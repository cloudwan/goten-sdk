package clipb

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cloudwan/goten-sdk/runtime/resource"
	"github.com/cloudwan/goten-sdk/runtime/utils"
)

type NameProvider struct {
	fromFlags   map[string]*string
	fromContext map[string]string
}

func NewNameProvider(segmentsFromContext map[string]string) *NameProvider {
	provider := &NameProvider{
		fromFlags:   make(map[string]*string),
		fromContext: make(map[string]string),
	}
	for key, value := range segmentsFromContext {
		provider.fromContext[strcase.ToKebab(key)] = value
	}
	return provider
}

func (p *NameProvider) RegisterFlagsFromRequest(flags *pflag.FlagSet, protoMsg proto.Message) {
	var registerRefsFunc func(flags *pflag.FlagSet, protoMsg proto.Message, populatedNames map[string]struct{})

	registerRefsFunc = func(flags *pflag.FlagSet, protoMsg proto.Message, populatedNames map[string]struct{}) {
		fields := protoMsg.ProtoReflect().Descriptor().Fields()
		for i := 0; i < fields.Len(); i++ {
			fieldDesc := fields.Get(i)
			if fieldDesc.IsList() || fieldDesc.IsMap() {
				continue
			} else if fieldDesc.Kind() == preflect.MessageKind || fieldDesc.Kind() == preflect.GroupKind {
				innerProtoMsg := protoMsg.ProtoReflect().NewField(fieldDesc).Message().Interface()
				registerRefsFunc(flags, innerProtoMsg, populatedNames)
			} else if fieldDesc.Kind() == preflect.StringKind {
				var fieldType reflect.StructField
				if fieldDesc.ContainingOneof() != nil {
					fieldType, _ = utils.GetFieldTypeForOneOf(protoMsg, fieldDesc)
				} else {
					fieldType, _ = reflect.TypeOf(protoMsg).Elem().FieldByName(strcase.ToCamel(string(fieldDesc.Name())))
				}
				defaultValue := reflect.Zero(fieldType.Type).Interface()
				if asName, ok := defaultValue.(resource.Name); ok {
					for idFieldName := range asName.GetIdParts() {
						resName := strings.TrimSuffix(idFieldName, "Id")
						if _, set := populatedNames[resName]; set {
							continue
						}
						populatedNames[resName] = struct{}{}
						kebabCase := strcase.ToKebab(resName)
						camelCase := strcase.ToCamel(resName)
						value := p.fromFlags[kebabCase]
						if value == nil {
							value = new(string)
							p.fromFlags[kebabCase] = value
						}
						if flags.Lookup(kebabCase) == nil {
							flags.StringVar(value, kebabCase, "", camelCase)
						} else {
							flags.StringVar(value, fmt.Sprintf("%s-name", kebabCase), "", camelCase)
						}
					}
				}
			}
		}
	}
	registerRefsFunc(flags, protoMsg, make(map[string]struct{}))
}

type pathSelection struct {
	segments         resource.NameSegments
	specifiedByFlags int
	specifiedByCtx   int
}

func (p *NameProvider) FillName(name resource.Name) {
	// dont attempt to override
	if name.IsSpecified() {
		return
	}

	nameDescriptor := name.GetResourceDescriptor().GetNameDescriptor()

	// (tricky) if GetIdParts does not return ID field name for resource, we are dealing with parent name
	_, hasIdFieldName := name.GetIdParts()[nameDescriptor.GetIdFieldName()]
	isParentName := !hasIdFieldName

	// Prefer name provided by flags over name context. Then, prefer longer name patterns.
	// Finally, prefer shorter paths.
	//
	// Flags are assumed as overriding values in the context, they can only complement each other.
	// Also, if name patterns are like "lastName" + "firstName/lastName", then if user
	// provides both arguments, make sure we dont default on the first shorten pattern.
	// We may mix name pattern provided from context and flags, if we cant construct
	// full name from flags only, unless mixed name contains name provided by flags only as suffix.
	//
	// If user provided "all" segments from multiple paths at once by flags (for example, lets say
	// possible patterns are "/ancestor1/..." AND "/ancestor2/...": and user provided both...),
	// then we just prefer longer then random path (which is fine, since user provided
	// ambiguous input anyway).
	//
	// Undefined segments populate with wildcards.

	var bestPath *pathSelection
	currentIds := name.GetIdParts()
	for _, pattern := range nameDescriptor.GetNamePatterns() {
		if name.GetPattern() != "" && name.GetPattern() != pattern {
			continue
		}

		path := &pathSelection{}
		if isParentName {
			path.segments = make([]resource.NameSegment, pattern.SegmentsCount()-1)
		} else {
			path.segments = make([]resource.NameSegment, pattern.SegmentsCount())
		}
		segmentPatterns := pattern.SegmentPatterns()[:len(path.segments)]

		for i, segmentPattern := range segmentPatterns {
			if valueFromName := currentIds[strcase.ToLowerCamel(segmentPattern.IdFieldName())]; valueFromName != "" {
				path.segments[i] = segmentPattern.MakeSegment(valueFromName)
				continue
			}
			segmentKey := strcase.ToKebab(segmentPattern.SingularLowerJson)
			valueFromFlag := p.fromFlags[segmentKey]
			if valueFromFlag != nil && *valueFromFlag != "" {
				path.specifiedByFlags++
				path.segments[i] = segmentPattern.MakeSegment(*valueFromFlag)
			} else {
				if valueFromCtx, isSet := p.fromContext[segmentKey]; isSet {
					path.specifiedByCtx++
					path.segments[i] = segmentPattern.MakeSegment(valueFromCtx)
				} else {
					path.segments[i] = segmentPattern.MakeSegment(resource.WildcardId)
				}
			}
		}
		if bestPath == nil || bestPath.specifiedByFlags < path.specifiedByFlags || (bestPath.specifiedByFlags == path.specifiedByFlags && (
				bestPath.specifiedByCtx < path.specifiedByCtx || bestPath.specifiedByCtx == path.specifiedByCtx && len(bestPath.segments) > len(path.segments))) {
			bestPath = path
		}
	}

	if bestPath == nil {
		panic(fmt.Errorf("could not provide %s", name))
	}
	if err := name.SetFromSegments(bestPath.segments); err != nil {
		panic(err)
	}
}
