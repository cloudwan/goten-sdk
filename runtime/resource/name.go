package resource

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"

	preflect "google.golang.org/protobuf/reflect/protoreflect"
)

const (
	WildcardId = "-"
)

// Name is an interface for goten name and reference goten types (including parent variants).
type Name interface {
	fmt.Stringer
	preflect.ProtoStringer

	// SetFromSegments modifies ids and pattern using given segments
	SetFromSegments(segments NameSegments) error

	// GotenEqual tells if other object is of same type and
	// is equal to current Name.
	GotenEqual(other interface{}) bool

	// Matches tells if other object is an object of
	// same type and is at least as specific as
	// current Name.
	Matches(other interface{}) bool

	// IsSpecified informs if Name has active pattern
	IsSpecified() bool

	// IsFullyQualified informs if Name has all id segments filled according
	// to it's active pattern. Non-Fully Qualified Names
	// can be used as path filters
	IsFullyQualified() bool

	// FullyQualifiedName gives resource name with service identifier
	// If IsFullyQualified is not true, then error will be
	// returned.
	// Format is: "//($SERVICE_NAME)/($RESOURCE_NAME), where $RESOURCE_NAME is equal to output of method String()
	FullyQualifiedName() (string, error)

	// AsReference<T>() *T_Reference

	// AsRawReference makes a new reference out of this name
	AsRawReference() Reference

	// GetResourceDescriptor returns associated with resource
	GetResourceDescriptor() Descriptor

	// GetPattern returns active pattern
	GetPattern() NamePattern

	// GetIdParts returns IDs for all possible name segments.
	// If name object does not contain some segment (is not fully specified OR segment is not present in
	// the active pattern), then value for given key is an empty string.
	GetIdParts() map[string]string

	// GetSegments returns list of segments according to the active name pattern
	GetSegments() NameSegments
}

// NamePattern represents structure of a resource name
type NamePattern string

func (np NamePattern) String() string {
	return string(np)
}

func (np NamePattern) SegmentsCount() int {
	return (strings.Count(string(np), "/") + 1) / 2
}

func (np NamePattern) SegmentPatterns() NameSegmentPatterns {
	items := strings.Split(string(np), "/")
	segmentPatterns := make(NameSegmentPatterns, 0, len(items) / 2)
	for i := 0; i < len(items); i += 2 {
		segmentPatterns = append(segmentPatterns, NameSegmentPattern{
			CollectionLowerJson: items[i],
			SingularLowerJson:   strcase.ToLowerCamel(items[i+1][1:len(items[i+1])-1]),
		})
	}
	return segmentPatterns
}

func (np NamePattern) IsPrefixOf(other NamePattern) bool {
	return strings.HasPrefix(string(other), string(np))
}

// NameSegment represents single segment within resource name. Consists of
// collection name and identifier.
type NameSegment struct {
	CollectionLowerJson string
	Id                  string
}

func (ns NameSegment) String() string {
	return fmt.Sprintf("%s/%s", ns.CollectionLowerJson, ns.Id)
}

type NameSegments []NameSegment

func (ns NameSegments) String() string {
	items := make([]string, 0, len(ns))
	for _, segment := range ns {
		items = append(items, segment.CollectionLowerJson)
		items = append(items, segment.Id)
	}
	return strings.Join(items, "/")
}

type NameSegmentPattern struct {
	CollectionLowerJson string
	SingularLowerJson   string
}

func (ns NameSegmentPattern) String() string {
	return fmt.Sprintf("%s/{%s}", ns.CollectionLowerJson, strcase.ToSnake(ns.SingularLowerJson))
}

func (ns NameSegmentPattern) MakeSegment(idValue string) NameSegment {
	return NameSegment{
		CollectionLowerJson: ns.CollectionLowerJson,
		Id:                  fmt.Sprintf("%s", idValue),
	}
}

func (ns NameSegmentPattern) IdFieldName() string {
	return fmt.Sprintf("%sId", ns.SingularLowerJson)
}

type NameSegmentPatterns []NameSegmentPattern

func (nsp NameSegmentPatterns) String() string {
	items := make([]string, 0, len(nsp))
	for _, segmentPattern := range nsp {
		items = append(items, segmentPattern.String())
	}
	return strings.Join(items, "/")
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

var idCharsUnicode = regexp.MustCompile(`[\p{L}\d]+`)
var idChars = regexp.MustCompile(`[\dA-Za-z]+`)
var idremalphabet = []rune("0123456789abcdefghijkmnopqrstuvwxyz")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := 0; i < n; i++ {
		b[i] = idremalphabet[rand.Intn(len(idremalphabet))]
	}
	return string(b)
}

func GenerateResourceId(res Resource) string {
	var segments = make([]string, 0)
	var idPrefix string
	if dres, ok := res.(DisplayableResource); ok && dres.GetDisplayName() != "" {
		words := idChars.FindAllString(strings.ToLower(dres.GetDisplayName()), -1)
		idPrefix = strings.Join(words, "-")
		idPrefix = idPrefix[:min(24, len(idPrefix))]
	}
	if idPrefix != "" {
		segments = append(segments, idPrefix)
	}
	rem := randStringRunes(max(6, 16-len(idPrefix)))
	if rem != "" {
		segments = append(segments, rem)
	}

	return strings.Join(segments, "-")
}
