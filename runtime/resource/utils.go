package resource

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	mrann "github.com/cloudwan/goten-sdk/annotations/multi_region"
	resourceann "github.com/cloudwan/goten-sdk/annotations/resource"
	"github.com/cloudwan/goten-sdk/runtime/object"
	mrp "github.com/cloudwan/goten-sdk/types/multi_region_policy"
)

const (
	RegionScopeAttribute = "goten.annotations/Region"
)

func IsRegionalResource(descriptor Descriptor) bool {
	msgOpts := descriptor.NewResource().ProtoReflect().Descriptor().Options().(*descriptorpb.MessageOptions)
	resSpec := proto.GetExtension(msgOpts, resourceann.E_Resource).(*resourceann.ResourceSpec)
	for _, attribute := range resSpec.GetScopeAttributes() {
		if attribute == RegionScopeAttribute {
			return true
		}
	}

	// TODO: We dont support maybe-regional resources depending on parents... if we dont
	// have RegionScopeAttribute, then ALL parents must return same value of IsRegionalResource
	// We need probably validation on api-skeleton level.
	for _, parentDesc := range descriptor.GetParentResDescriptors() {
		if IsRegionalResource(parentDesc) {
			return true
		}
	}
	return false
}

func GetSyncingOpts(descriptor Descriptor) *mrann.SyncingOpts {
	msgDes := descriptor.NewResource().ProtoReflect().Descriptor()
	msgOpts := msgDes.Options().(*descriptorpb.MessageOptions)
	syncingOpts := proto.GetExtension(msgOpts, mrann.E_SyncingOpts).(*mrann.SyncingOpts)
	return syncingOpts
}

func FindMultiRegionPolicyFieldPath(resDescriptor Descriptor) object.FieldPath {
	var nilPolicy *mrp.MultiRegionPolicy
	searchedMsgDes := nilPolicy.ProtoReflect().Descriptor()

	resourceProtoDesc := resDescriptor.NewResource().ProtoReflect().Descriptor()
	fieldsStack := make([][]preflect.FieldDescriptor, 0)
	msgsStack := make([]preflect.MessageDescriptor, 0)

	msgsStack = append(msgsStack, resourceProtoDesc)
	fieldsStack = append(fieldsStack, []preflect.FieldDescriptor{})
	var visitFunc func(idx int) bool

	visited := map[string]bool{}
	visitFunc = func(idx int) bool {
		currentMsg := msgsStack[idx]
		currentParentFields := fieldsStack[idx]

		if currentMsg.FullName() == searchedMsgDes.FullName() {
			return true
		}

		fieldDescriptors := currentMsg.Fields()
		for i := 0; i < fieldDescriptors.Len(); i++ {
			fieldDes := fieldDescriptors.Get(i)
			if fieldDes.Kind() == preflect.MessageKind && visited[string(fieldDes.Message().FullName())] == false {
				visited[string(fieldDes.Message().FullName())] = true
				msgsStack = append(msgsStack, fieldDes.Message())
				newParentFields := make([]preflect.FieldDescriptor, 0)
				newParentFields = append(newParentFields, currentParentFields...)
				newParentFields = append(newParentFields, fieldDes)
				fieldsStack = append(fieldsStack, newParentFields)
			}
		}
		return false
	}

	for i := 0; i < len(msgsStack); i++ {
		if visitFunc(i) {
			rawFieldPath := ""

			for _, fieldDes := range fieldsStack[i] {
				if rawFieldPath != "" {
					rawFieldPath += "."
				}
				rawFieldPath += string(fieldDes.Name())
			}
			gotenFp, err := resDescriptor.ParseFieldPath(rawFieldPath)
			if err != nil {
				panic(fmt.Sprintf("Failed to parse field path %s: %s", rawFieldPath, err))
			}
			return gotenFp
		}
	}
	return nil
}
