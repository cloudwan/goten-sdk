package resource

import (
	"sort"

	"github.com/cloudwan/goten-sdk/runtime/object"
	"github.com/cloudwan/goten-sdk/runtime/strcase"
)

// Descriptor allows writing code operating on resources without knowing exact type.
// It can be used to create Resource instance, all derivative types, access information
// about name type.
type Descriptor interface {
	NewResource() Resource
	NewResourceChange() ResourceChange
	NewResourceName() Name
	NewResourceCursor() Cursor
	NewResourceFilter() Filter
	NewResourceOrderBy() OrderBy
	NewResourcePager() PagerQuery
	NewResourceFieldMask() object.FieldMask

	NewGetQuery() GetQuery
	NewListQuery() ListQuery
	NewSearchQuery() SearchQuery
	NewWatchQuery() WatchQuery
	NewQueryResultSnapshot() QueryResultSnapshot
	NewQueryResultChange() QueryResultChange

	NewResourceList(size, reserved int) ResourceList
	NewResourceChangeList(size, reserved int) ResourceChangeList
	NewNameList(size, reserved int) NameList
	NewReferenceList(size, reserved int) ReferenceList
	NewParentNameList(size, reserved int) ParentNameList
	NewParentReferenceList(size, reserved int) ParentReferenceList
	NewResourceMap(reserved int) ResourceMap
	NewResourceChangeMap(reserved int) ResourceChangeMap

	GetResourceTypeName() *TypeName
	GetNameDescriptor() *NameDescriptor
	CanBeParentless() bool
	GetParentResDescriptors() []Descriptor

	SupportsMetadata() bool
	SupportsDbConstraints() bool

	// ParseFieldPath returns a field path matching given string. If does not exist, error
	// is returned and returned interface is None.
	ParseFieldPath(raw string) (object.FieldPath, error)
	ParseResourceName(raw string) (Name, error)
}

type NameDescriptor struct {
	fieldPath          object.FieldPath
	patternFieldName   string
	idFieldName        string
	parentIdFieldNames []string
	namePatterns       []NamePattern
}

func NewNameDescriptor(fp object.FieldPath, patternFieldName, idFieldName string,
	parentIdFieldNames []string, namePatterns []NamePattern) *NameDescriptor {
	return &NameDescriptor{
		fieldPath:          fp,
		patternFieldName:   patternFieldName,
		idFieldName:        idFieldName,
		parentIdFieldNames: parentIdFieldNames,
		namePatterns:       namePatterns,
	}
}

func (d *NameDescriptor) GetFieldPath() object.FieldPath {
	return d.fieldPath
}

func (d *NameDescriptor) GetPatternFieldName() string {
	return d.patternFieldName
}

func (d *NameDescriptor) GetIdFieldName() string {
	return d.idFieldName
}

func (d *NameDescriptor) GetParentIdFieldNames() []string {
	return d.parentIdFieldNames
}

func (d *NameDescriptor) GetAllSegmentIdFieldNames() []string {
	res := d.parentIdFieldNames
	return append(res, d.idFieldName)
}

func (d *NameDescriptor) GetNamePatterns() []NamePattern {
	return d.namePatterns
}

type TypeName struct {
	singular string
	plural   string
	domain   string
	version  string
}

func NewTypeName(singular, plural, domain, version string) *TypeName {
	return &TypeName{singular: singular, plural: plural, domain: domain, version: version}
}

func (rtn *TypeName) Singular() string {
	return rtn.singular
}

func (rtn *TypeName) Plural() string {
	return rtn.plural
}

func (rtn *TypeName) Version() string {
	return rtn.version
}

func (rtn *TypeName) JSONSingular() string {
	return strcase.ToLowerCamel(rtn.singular)
}

func (rtn *TypeName) JSONPlural() string {
	return strcase.ToLowerCamel(rtn.plural)
}

func (rtn *TypeName) FullyQualifiedTypeName() string {
	return rtn.domain + "/" + rtn.singular
}

func (rtn *TypeName) ServiceDomain() string {
	return rtn.domain
}

func SortedResourceIdRefNameSegments(descriptor Descriptor) []string {
	nameToSegment := make(map[string]NameSegmentPattern)
	havingOnTheLeftSide := make(map[string]map[string]NameSegmentPattern)
	havingOnTheRightSide := make(map[string]map[string]NameSegmentPattern)

	for _, namePattern := range descriptor.GetNameDescriptor().GetNamePatterns() {
		segmentPatterns := namePattern.SegmentPatterns()
		for i, segmentPattern := range segmentPatterns {
			collection := segmentPattern.CollectionLowerJson
			nameToSegment[collection] = segmentPattern
			if havingOnTheLeftSide[collection] == nil {
				havingOnTheLeftSide[collection] = make(map[string]NameSegmentPattern)
			}
			if havingOnTheRightSide[collection] == nil {
				havingOnTheRightSide[collection] = make(map[string]NameSegmentPattern)
			}
			onTheLeftSide := segmentPatterns[0:i]
			for _, segmentOnTheLeftSide := range onTheLeftSide {
				leftCollection := segmentOnTheLeftSide.CollectionLowerJson
				havingOnTheRightSide[leftCollection][collection] = segmentPattern
				havingOnTheLeftSide[collection][leftCollection] = segmentOnTheLeftSide
			}
		}
	}

	result := make([]string, 0)
	for len(havingOnTheLeftSide) > 0 {
		toAppend := make([]NameSegmentPattern, 0)
		for segmentName, remainingLeftNeighbours := range havingOnTheLeftSide {
			if len(remainingLeftNeighbours) == 0 {
				toAppend = append(toAppend, nameToSegment[segmentName])
			}
		}
		sort.Slice(toAppend, func(i, j int) bool {
			return toAppend[i].IdFieldName() < toAppend[j].IdFieldName()
		})
		for _, nameSegment := range toAppend {
			delete(havingOnTheLeftSide, nameSegment.CollectionLowerJson)
			neighboursOnRight := havingOnTheRightSide[nameSegment.CollectionLowerJson]
			for rightNeighbourName := range neighboursOnRight {
				delete(havingOnTheLeftSide[rightNeighbourName], nameSegment.CollectionLowerJson)
			}
			result = append(result, nameSegment.IdFieldName())
		}
	}
	return result
}
