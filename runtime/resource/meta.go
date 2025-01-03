package resource

import (
	"fmt"
	"google.golang.org/protobuf/proto"

	"github.com/cloudwan/goten-sdk/types/meta"
)

func FindMetaOwnerByDescriptor(res Resource, descriptor Descriptor) (mref *meta.OwnerReference) {
	for _, oref := range res.GetMetadata().GetOwnerReferences() {
		if oref.GetVersion() == descriptor.GetResourceTypeName().Version() &&
			oref.GetKind() == descriptor.GetResourceTypeName().FullyQualifiedTypeName() {
			mref = oref
			break
		}
	}
	return
}

type OwnerRefFlag func(o *meta.OwnerReference)

func WithController() OwnerRefFlag {
	return func(o *meta.OwnerReference) {
		o.Controller = true
	}
}

func WithUnsetOnDelete() OwnerRefFlag {
	return func(o *meta.OwnerReference) {
		o.UnsetOnDelete = true
	}
}

// TODO: Unsupported
/*
func WithBlockDeletion() OwnerRefFlag {
	return func(o *OwnerReference) {
		o.BlockOwnerDeletion = true
	}
}
*/

func WithRequiresOwnerReference() OwnerRefFlag {
	return func(o *meta.OwnerReference) {
		o.RequiresOwnerReference = true
	}
}

func MakeMetaOwnerReference(owner Resource, flags ...OwnerRefFlag) *meta.OwnerReference {
	ownerDescriptor := owner.GetResourceDescriptor()
	ownerRef := &meta.OwnerReference{
		Kind:    ownerDescriptor.GetResourceTypeName().FullyQualifiedTypeName(),
		Version: ownerDescriptor.GetResourceTypeName().Version(),
		Name:    owner.GetRawName().String(),
		Region:  owner.GetMetadata().GetSyncing().GetOwningRegion(),
	}
	if ownerRef.GetRegion() == "" {
		if IsRegionalResource(ownerDescriptor) {
			ownerRef.Region = owner.GetRawName().GetIdParts()["regionId"]
		}
	}
	if ownerRef.GetRegion() == "" {
		panic(fmt.Errorf(
			"MakeMetaOwnerReference must be called with a resource containing region ID in either "+
				"metadata.syncing.owning_region field path, or with region ID segment in the name field. Got resource "+
				"with name %s, and metadata.syncing.owning_region is unset",
			owner.GetRawName()))
	}
	for _, flag := range flags {
		flag(ownerRef)
	}
	return ownerRef
}

func MakeMetaOwnerReferenceFromNameAndRegion(ownerName Name, ownerRegionId string, flags ...OwnerRefFlag) *meta.OwnerReference {
	ownerDescriptor := ownerName.GetResourceDescriptor()
	ownerRef := &meta.OwnerReference{
		Kind:    ownerDescriptor.GetResourceTypeName().FullyQualifiedTypeName(),
		Version: ownerDescriptor.GetResourceTypeName().Version(),
		Name:    ownerName.String(),
		Region:  ownerRegionId,
	}
	for _, flag := range flags {
		flag(ownerRef)
	}
	return ownerRef
}

func AddUniqueMetaOwnerReference(res Resource, ownerRef *meta.OwnerReference) bool {
	m := res.EnsureMetadata()
	for _, current := range m.GetOwnerReferences() {
		if current.GetName() == ownerRef.GetName() {
			proto.Merge(current, ownerRef)
			return false
		}
	}
	m.OwnerReferences = append(m.OwnerReferences, ownerRef)
	return true
}
