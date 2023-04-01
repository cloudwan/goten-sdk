package service_model

import (
	"fmt"
	"sort"
	"strings"

	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	ann "github.com/cloudwan/goten-sdk/annotations/goten"
	resourceann "github.com/cloudwan/goten-sdk/annotations/resource"
	"github.com/cloudwan/goten-sdk/runtime/client"
	"github.com/cloudwan/goten-sdk/runtime/object"
	"github.com/cloudwan/goten-sdk/runtime/resource"
)

type ServiceModel struct {
	svcDescriptor    client.ServiceDescriptor
	orderedResources []*ResourceModel
}

func MakeServiceModel(svcDescriptor client.ServiceDescriptor) *ServiceModel {
	svcModel := &ServiceModel{
		svcDescriptor: svcDescriptor,
	}
	resModels := make([]*ResourceModel, 0)
	for _, resDesc := range svcDescriptor.AllResourceDescriptors() {
		resModels = append(resModels, makeResourceModel(resDesc))
	}
	populateBackRefDeletionTriggers(resModels)
	svcModel.orderedResources = sortResourcesMostBasicFirst(resModels)
	return svcModel
}

func (m *ServiceModel) GetSvcDescriptor() client.ServiceDescriptor {
	return m.svcDescriptor
}

func (m *ServiceModel) GetResourceModels() []*ResourceModel {
	return m.orderedResources
}

type ResourceModel struct {
	descriptor              resource.Descriptor
	gotenSpec               *resourceann.ResourceSpec
	refInfos                []ReferenceInfo
	backRefsSyncDelHandlers []*BackRefDeletionHandler
}

func (m *ResourceModel) GetDescriptor() resource.Descriptor {
	return m.descriptor
}

func (m *ResourceModel) GetGotenSpec() *resourceann.ResourceSpec {
	return m.gotenSpec
}

func (m *ResourceModel) GetReferenceInfos() []ReferenceInfo {
	return m.refInfos
}

func (m *ResourceModel) GetSynchronousBackRefsDeletionHandlers() []*BackRefDeletionHandler {
	return m.backRefsSyncDelHandlers
}

type refFieldPathItemType int

const (
	refFieldPathItemTypeRef    refFieldPathItemType = 0
	refFieldPathItemTypeRefRep refFieldPathItemType = 1
	refFieldPathItemTypeObj    refFieldPathItemType = 2
	refFieldPathItemTypeObjRep refFieldPathItemType = 3
)

type ReferenceInfo struct {
	fieldPath          object.FieldPath
	onTargetDeleted    ann.OnReferencedObjectDeleted
	targetResName      string
	fieldPathItemTypes []refFieldPathItemType
}

func (ri *ReferenceInfo) GetFieldPath() object.FieldPath {
	return ri.fieldPath
}

func (ri *ReferenceInfo) GetOnTargetDeleted() ann.OnReferencedObjectDeleted {
	return ri.onTargetDeleted
}

func (ri *ReferenceInfo) GetTargetResName() string {
	return ri.targetResName
}

func (ri *ReferenceInfo) IsRepeated() bool {
	return ri.fieldPathItemTypes[len(ri.fieldPathItemTypes)-1] == refFieldPathItemTypeRefRep
}

func (ri *ReferenceInfo) HasRepeatedItemOnPath() bool {
	for _, item := range ri.fieldPathItemTypes {
		if item == refFieldPathItemTypeRefRep || item == refFieldPathItemTypeObjRep {
			return true
		}
	}
	return false
}

type BackRefDeletionHandler struct {
	unsetType         bool
	cascadeDeleteType bool

	backRefDescriptor  resource.Descriptor
	backRefFp          object.FieldPath
	fieldPathItemTypes []refFieldPathItemType
}

func (h *BackRefDeletionHandler) GetBackRefDescriptor() resource.Descriptor {
	return h.backRefDescriptor
}

func (h *BackRefDeletionHandler) GetBackRefFp() object.FieldPath {
	return h.backRefFp
}

func (h *BackRefDeletionHandler) IsUnset() bool {
	return h.unsetType
}

func (h *BackRefDeletionHandler) IsCascadeDelete() bool {
	return h.cascadeDeleteType
}

func (h *BackRefDeletionHandler) BackRefIsRepeated() bool {
	return h.fieldPathItemTypes[len(h.fieldPathItemTypes)-1] == refFieldPathItemTypeRefRep
}

func (h *BackRefDeletionHandler) HasRepeatedItemOnPath() bool {
	for _, item := range h.fieldPathItemTypes {
		if item == refFieldPathItemTypeRefRep || item == refFieldPathItemTypeObjRep {
			return true
		}
	}
	return false
}

type tmpRefInfo struct {
	raw                string
	fieldPathItemTypes []refFieldPathItemType
	targetResName      string
	onDelBehavior      ann.OnReferencedObjectDeleted
}

func makeResourceModel(descriptor resource.Descriptor) *ResourceModel {
	schema := &ResourceModel{descriptor: descriptor}
	msgDescriptor := descriptor.NewResource().ProtoReflect().Descriptor()

	knownRefsForMsg := map[string][]tmpRefInfo{}
	var collectRefPaths func(msgDes preflect.MessageDescriptor, prefix string, typesSoFar []refFieldPathItemType) []tmpRefInfo

	collectRefPaths = func(msgDes preflect.MessageDescriptor, prefix string, typesSoFar []refFieldPathItemType) []tmpRefInfo {
		fields := msgDes.Fields()
		objectRefInfos := make([]tmpRefInfo, 0)
		for i := 0; i < fields.Len(); i++ {
			fd := fields.Get(i)
			if fd.IsMap() {
				// TODO: Maps!
				continue
			}

			if fd.Kind() == preflect.GroupKind || fd.Kind() == preflect.MessageKind {
				subMsgDes := fd.Message()
				objTypeName := string(subMsgDes.FullName())
				itemFieldType := refFieldPathItemTypeObj
				if fd.IsList() {
					itemFieldType = refFieldPathItemTypeObjRep
				}
				fpItemTypesForCurrentPath := make([]refFieldPathItemType, 0, len(typesSoFar)+1)
				fpItemTypesForCurrentPath = append(fpItemTypesForCurrentPath, typesSoFar...)
				fpItemTypesForCurrentPath = append(fpItemTypesForCurrentPath, itemFieldType)

				var ok bool
				var refsForMsg []tmpRefInfo
				if refsForMsg, ok = knownRefsForMsg[string(subMsgDes.FullName())]; !ok {
					knownRefsForMsg[objTypeName] = make([]tmpRefInfo, 0)
					refsForMsg = collectRefPaths(subMsgDes, fmt.Sprintf("%s%s.", prefix, fd.Name()), fpItemTypesForCurrentPath)
					knownRefsForMsg[objTypeName] = refsForMsg
				} else {
					for _, refInfo := range refsForMsg {
						rawProtoFieldPath := fmt.Sprintf("%s%s.%s", prefix, fd.Name(), refInfo.raw)
						fpItemTypesForThisRef := make([]refFieldPathItemType, 0, len(fpItemTypesForCurrentPath)+len(refInfo.fieldPathItemTypes))
						fpItemTypesForThisRef = append(fpItemTypesForThisRef, fpItemTypesForCurrentPath...)
						fpItemTypesForThisRef = append(fpItemTypesForThisRef, refInfo.fieldPathItemTypes...)

						fp, err := descriptor.ParseFieldPath(rawProtoFieldPath)
						if err == nil {
							schema.refInfos = append(schema.refInfos, ReferenceInfo{
								fieldPath:          fp,
								onTargetDeleted:    refInfo.onDelBehavior,
								targetResName:      refInfo.targetResName,
								fieldPathItemTypes: fpItemTypesForThisRef,
							})
						} else {
							panic(fmt.Errorf("error parsing field path %s for resource %s: %s",
								rawProtoFieldPath, descriptor.GetResourceTypeName().FullyQualifiedTypeName(), err))
						}
					}
				}
				for _, refInfo := range refsForMsg {
					fpItemTypesForThisObject := make([]refFieldPathItemType, 0, len(refInfo.fieldPathItemTypes)+1)
					fpItemTypesForThisObject = append(fpItemTypesForThisObject, itemFieldType)
					fpItemTypesForThisObject = append(fpItemTypesForThisObject, refInfo.fieldPathItemTypes...)
					objectRefInfos = append(objectRefInfos, tmpRefInfo{
						raw:                fmt.Sprintf("%s.%s", fd.Name(), refInfo.raw),
						fieldPathItemTypes: fpItemTypesForThisObject,
						targetResName:      refInfo.targetResName,
						onDelBehavior:      refInfo.onDelBehavior,
					})
				}
			} else if fd.Kind() == preflect.StringKind {
				fieldOpts := fd.Options().(*descriptorpb.FieldOptions)
				fieldTypeSpec := proto.GetExtension(fieldOpts, ann.E_Type).(*ann.FieldType)
				if fieldTypeSpec == nil {
					continue
				}

				if refFt, ok := fieldTypeSpec.Type.(*ann.FieldType_Reference); ok {
					itemFieldType := refFieldPathItemTypeRef
					if fd.IsList() {
						itemFieldType = refFieldPathItemTypeRefRep
					}
					fpItemTypesForThisRef := make([]refFieldPathItemType, 0, len(typesSoFar)+1)
					fpItemTypesForThisRef = append(fpItemTypesForThisRef, typesSoFar...)
					fpItemTypesForThisRef = append(fpItemTypesForThisRef, itemFieldType)

					rawProtoFieldPath := fmt.Sprintf("%s%s", prefix, fd.Name())
					fp, err := descriptor.ParseFieldPath(rawProtoFieldPath)
					if err == nil {
						objectRefInfos = append(objectRefInfos, tmpRefInfo{
							raw:                string(fd.Name()),
							fieldPathItemTypes: []refFieldPathItemType{itemFieldType},
							targetResName:      refFt.Reference.GetResource(),
							onDelBehavior:      refFt.Reference.GetTargetDeleteBehavior(),
						})
						schema.refInfos = append(schema.refInfos, ReferenceInfo{
							fieldPath:          fp,
							onTargetDeleted:    refFt.Reference.GetTargetDeleteBehavior(),
							targetResName:      refFt.Reference.GetResource(),
							fieldPathItemTypes: fpItemTypesForThisRef,
						})
					} else {
						panic(fmt.Errorf("error parsing field path %s for resource %s: %s",
							rawProtoFieldPath, descriptor.GetResourceTypeName().FullyQualifiedTypeName(), err))
					}
				}
			}
		}
		return objectRefInfos
	}

	collectRefPaths(msgDescriptor, "", nil)
	msgOpts := msgDescriptor.Options().(*descriptorpb.MessageOptions)
	resSpec := proto.GetExtension(msgOpts, resourceann.E_Spec).(*resourceann.ResourceSpec)
	schema.gotenSpec = resSpec
	return schema
}

func populateBackRefDeletionTriggers(resModels []*ResourceModel) {
	for _, resModel := range resModels {
		if resModel.gotenSpec.GetOnParentDeletedBehavior() == ann.OnReferencedObjectDeleted_CASCADE_DELETE {
			for _, parentName := range resModel.gotenSpec.GetParents() {
				parentModel := findResByName(resModels, parentName)
				if parentModel == nil {
					panic("Could not find local definition for parent " + parentName)
				}
				handler := &BackRefDeletionHandler{
					cascadeDeleteType:  true,
					unsetType:          false,
					backRefDescriptor:  resModel.descriptor,
					backRefFp:          resModel.descriptor.GetNameDescriptor().GetFieldPath(),
					fieldPathItemTypes: []refFieldPathItemType{refFieldPathItemTypeRef},
				}
				parentModel.backRefsSyncDelHandlers = append(parentModel.backRefsSyncDelHandlers, handler)
			}
		}

		for _, refInfo := range resModel.refInfos {
			if refInfo.onTargetDeleted == ann.OnReferencedObjectDeleted_CASCADE_DELETE ||
				refInfo.onTargetDeleted == ann.OnReferencedObjectDeleted_UNSET {
				refModel := findResByName(resModels, refInfo.targetResName)
				if refModel == nil {
					panic("Could not find local definition for resource " + refInfo.targetResName)
				}
				handler := &BackRefDeletionHandler{
					cascadeDeleteType:  refInfo.onTargetDeleted == ann.OnReferencedObjectDeleted_CASCADE_DELETE,
					unsetType:          refInfo.onTargetDeleted == ann.OnReferencedObjectDeleted_UNSET,
					backRefDescriptor:  resModel.descriptor,
					backRefFp:          refInfo.fieldPath,
					fieldPathItemTypes: refInfo.fieldPathItemTypes,
				}
				refModel.backRefsSyncDelHandlers = append(refModel.backRefsSyncDelHandlers, handler)
			}
		}
	}
}

func sortResourcesMostBasicFirst(resModels []*ResourceModel) []*ResourceModel {
	depsByRes := make(map[string]map[string]struct{})
	for _, model := range resModels {
		deps := make(map[string]struct{})
		for _, parentName := range model.gotenSpec.GetParents() {
			if depModel := findResByName(resModels, parentName); depModel != nil {
				deps[parentName] = struct{}{}
			}
		}
		for _, ref := range model.refInfos {
			if depModel := findResByName(resModels, ref.targetResName); depModel != nil {
				deps[ref.targetResName] = struct{}{}
			}
		}
		depsByRes[model.descriptor.GetResourceTypeName().Singular()] = deps
	}

	output := make([]*ResourceModel, 0)
	inOutput := make(map[string]struct{})
	for len(output) != len(resModels) {
		nextToOut := make([]*ResourceModel, 0)
		for _, model := range resModels {
			singularName := model.descriptor.GetResourceTypeName().Singular()
			if _, added := inOutput[singularName]; added {
				continue
			}
			canBeCandidate := true
			for depName := range depsByRes[singularName] {
				if depName == singularName {
					continue
				}
				if _, added := inOutput[depName]; !added {
					canBeCandidate = false
					break
				}
			}
			if canBeCandidate {
				nextToOut = append(nextToOut, model)
			}
		}
		sort.Slice(nextToOut, func(i, j int) bool {
			return nextToOut[i].descriptor.GetResourceTypeName().Singular() < nextToOut[j].descriptor.GetResourceTypeName().Singular()
		})
		for _, next := range nextToOut {
			inOutput[next.descriptor.GetResourceTypeName().Singular()] = struct{}{}
		}
		output = append(output, nextToOut...)
	}
	return output
}

func findResByName(allModels []*ResourceModel, name string) *ResourceModel {
	if !strings.Contains(name, "/") {
		for _, model := range allModels {
			if model.descriptor.GetResourceTypeName().Singular() == name {
				return model
			}
		}
	}
	return nil
}
