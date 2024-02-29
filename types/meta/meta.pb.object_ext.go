// Code generated by protoc-gen-goten-object
// File: goten/types/meta.proto
// DO NOT EDIT!!!

package meta

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// ensure the imports are used
var (
	_ = new(fmt.Stringer)
	_ = new(sort.Interface)

	_ = new(proto.Message)
	_ = googlefieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &timestamppb.Timestamp{}
)

func (o *Meta) GotenObjectExt() {}

func (o *Meta) MakeFullFieldMask() *Meta_FieldMask {
	return FullMeta_FieldMask()
}

func (o *Meta) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullMeta_FieldMask()
}

func (o *Meta) MakeDiffFieldMask(other *Meta) *Meta_FieldMask {
	if o == nil && other == nil {
		return &Meta_FieldMask{}
	}
	if o == nil || other == nil {
		return FullMeta_FieldMask()
	}

	res := &Meta_FieldMask{}
	if !proto.Equal(o.GetCreateTime(), other.GetCreateTime()) {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorCreateTime})
	}
	if !proto.Equal(o.GetUpdateTime(), other.GetUpdateTime()) {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorUpdateTime})
	}
	if !proto.Equal(o.GetDeleteTime(), other.GetDeleteTime()) {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorDeleteTime})
	}
	if o.GetUuid() != other.GetUuid() {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorUuid})
	}

	if len(o.GetTags()) == len(other.GetTags()) {
		for i, lValue := range o.GetTags() {
			rValue := other.GetTags()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorTags})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorTags})
	}

	if len(o.GetLabels()) == len(other.GetLabels()) {
		for i, lValue := range o.GetLabels() {
			rValue := other.GetLabels()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorLabels})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorLabels})
	}

	if len(o.GetAnnotations()) == len(other.GetAnnotations()) {
		for i, lValue := range o.GetAnnotations() {
			rValue := other.GetAnnotations()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorAnnotations})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorAnnotations})
	}
	if o.GetGeneration() != other.GetGeneration() {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorGeneration})
	}
	if o.GetResourceVersion() != other.GetResourceVersion() {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorResourceVersion})
	}

	if len(o.GetOwnerReferences()) == len(other.GetOwnerReferences()) {
		for i, lValue := range o.GetOwnerReferences() {
			rValue := other.GetOwnerReferences()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorOwnerReferences})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorOwnerReferences})
	}

	if len(o.GetShards()) == len(other.GetShards()) {
		for i, lValue := range o.GetShards() {
			rValue := other.GetShards()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorShards})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorShards})
	}
	{
		subMask := o.GetSyncing().MakeDiffFieldMask(other.GetSyncing())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorSyncing})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Meta_FieldSubPath{selector: Meta_FieldPathSelectorSyncing, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetLifecycle().MakeDiffFieldMask(other.GetLifecycle())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorLifecycle})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Meta_FieldSubPath{selector: Meta_FieldPathSelectorLifecycle, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetServices().MakeDiffFieldMask(other.GetServices())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Meta_FieldTerminalPath{selector: Meta_FieldPathSelectorServices})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Meta_FieldSubPath{selector: Meta_FieldPathSelectorServices, subPath: subpath})
			}
		}
	}
	return res
}

func (o *Meta) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Meta))
}

func (o *Meta) Clone() *Meta {
	if o == nil {
		return nil
	}
	result := &Meta{}
	result.CreateTime = proto.Clone(o.CreateTime).(*timestamppb.Timestamp)
	result.UpdateTime = proto.Clone(o.UpdateTime).(*timestamppb.Timestamp)
	result.DeleteTime = proto.Clone(o.DeleteTime).(*timestamppb.Timestamp)
	result.Uuid = o.Uuid
	result.Tags = make([]string, len(o.Tags))
	for i, sourceValue := range o.Tags {
		result.Tags[i] = sourceValue
	}
	result.Labels = map[string]string{}
	for key, sourceValue := range o.Labels {
		result.Labels[key] = sourceValue
	}
	result.Annotations = map[string]string{}
	for key, sourceValue := range o.Annotations {
		result.Annotations[key] = sourceValue
	}
	result.Generation = o.Generation
	result.ResourceVersion = o.ResourceVersion
	result.OwnerReferences = make([]*OwnerReference, len(o.OwnerReferences))
	for i, sourceValue := range o.OwnerReferences {
		result.OwnerReferences[i] = sourceValue.Clone()
	}
	result.Shards = map[string]int64{}
	for key, sourceValue := range o.Shards {
		result.Shards[key] = sourceValue
	}
	result.Syncing = o.Syncing.Clone()
	result.Lifecycle = o.Lifecycle.Clone()
	result.Services = o.Services.Clone()
	return result
}

func (o *Meta) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Meta) Merge(source *Meta) {
	if source.GetCreateTime() != nil {
		if o.CreateTime == nil {
			o.CreateTime = new(timestamppb.Timestamp)
		}
		proto.Merge(o.CreateTime, source.GetCreateTime())
	}
	if source.GetUpdateTime() != nil {
		if o.UpdateTime == nil {
			o.UpdateTime = new(timestamppb.Timestamp)
		}
		proto.Merge(o.UpdateTime, source.GetUpdateTime())
	}
	if source.GetDeleteTime() != nil {
		if o.DeleteTime == nil {
			o.DeleteTime = new(timestamppb.Timestamp)
		}
		proto.Merge(o.DeleteTime, source.GetDeleteTime())
	}
	o.Uuid = source.GetUuid()
	for _, sourceValue := range source.GetTags() {
		exists := false
		for _, currentValue := range o.Tags {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.Tags = append(o.Tags, newDstElement)
		}
	}

	if source.GetLabels() != nil {
		if o.Labels == nil {
			o.Labels = make(map[string]string, len(source.GetLabels()))
		}
		for key, sourceValue := range source.GetLabels() {
			o.Labels[key] = sourceValue
		}
	}
	if source.GetAnnotations() != nil {
		if o.Annotations == nil {
			o.Annotations = make(map[string]string, len(source.GetAnnotations()))
		}
		for key, sourceValue := range source.GetAnnotations() {
			o.Annotations[key] = sourceValue
		}
	}
	o.Generation = source.GetGeneration()
	o.ResourceVersion = source.GetResourceVersion()
	sourceOwnerReferences, origOwnerReferencesKeys := map[string]*OwnerReference{}, map[string]bool{}
	newOwnerReferences := make([]*OwnerReference, 0, len(o.OwnerReferences))
	for _, sourceValue := range source.GetOwnerReferences() {
		key := fmt.Sprintf("%s", sourceValue.GetName())
		sourceOwnerReferences[key] = sourceValue
	}
	for _, origValue := range o.OwnerReferences {
		key := fmt.Sprintf("%s", origValue.GetName())
		origOwnerReferencesKeys[key] = true
		sourceValue := sourceOwnerReferences[key]
		if sourceValue != nil {
			if origValue == nil {
				origValue = new(OwnerReference)
			}
			origValue.Merge(sourceValue)
		}
		newOwnerReferences = append(newOwnerReferences, origValue)
	}
	for key, sourceValue := range sourceOwnerReferences {
		if origOwnerReferencesKeys[key] == false {
			newOwnerReferences = append(newOwnerReferences, sourceValue.Clone())
		}
	}
	o.OwnerReferences = newOwnerReferences

	if source.GetShards() != nil {
		if o.Shards == nil {
			o.Shards = make(map[string]int64, len(source.GetShards()))
		}
		for key, sourceValue := range source.GetShards() {
			o.Shards[key] = sourceValue
		}
	}
	if source.GetSyncing() != nil {
		if o.Syncing == nil {
			o.Syncing = new(SyncingMeta)
		}
		o.Syncing.Merge(source.GetSyncing())
	}
	if source.GetLifecycle() != nil {
		if o.Lifecycle == nil {
			o.Lifecycle = new(Lifecycle)
		}
		o.Lifecycle.Merge(source.GetLifecycle())
	}
	if source.GetServices() != nil {
		if o.Services == nil {
			o.Services = new(ServicesInfo)
		}
		o.Services.Merge(source.GetServices())
	}
}

func (o *Meta) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Meta))
}

func (o *LabelSelector) GotenObjectExt() {}

func (o *LabelSelector) MakeFullFieldMask() *LabelSelector_FieldMask {
	return FullLabelSelector_FieldMask()
}

func (o *LabelSelector) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLabelSelector_FieldMask()
}

func (o *LabelSelector) MakeDiffFieldMask(other *LabelSelector) *LabelSelector_FieldMask {
	if o == nil && other == nil {
		return &LabelSelector_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLabelSelector_FieldMask()
	}

	res := &LabelSelector_FieldMask{}

	if len(o.GetMatchLabels()) == len(other.GetMatchLabels()) {
		for i, lValue := range o.GetMatchLabels() {
			rValue := other.GetMatchLabels()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &LabelSelector_FieldTerminalPath{selector: LabelSelector_FieldPathSelectorMatchLabels})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &LabelSelector_FieldTerminalPath{selector: LabelSelector_FieldPathSelectorMatchLabels})
	}

	if len(o.GetMatchExpressions()) == len(other.GetMatchExpressions()) {
		for i, lValue := range o.GetMatchExpressions() {
			rValue := other.GetMatchExpressions()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &LabelSelector_FieldTerminalPath{selector: LabelSelector_FieldPathSelectorMatchExpressions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &LabelSelector_FieldTerminalPath{selector: LabelSelector_FieldPathSelectorMatchExpressions})
	}
	return res
}

func (o *LabelSelector) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*LabelSelector))
}

func (o *LabelSelector) Clone() *LabelSelector {
	if o == nil {
		return nil
	}
	result := &LabelSelector{}
	result.MatchLabels = map[string]string{}
	for key, sourceValue := range o.MatchLabels {
		result.MatchLabels[key] = sourceValue
	}
	result.MatchExpressions = make([]*LabelSelectorRequirement, len(o.MatchExpressions))
	for i, sourceValue := range o.MatchExpressions {
		result.MatchExpressions[i] = sourceValue.Clone()
	}
	return result
}

func (o *LabelSelector) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *LabelSelector) Merge(source *LabelSelector) {
	if source.GetMatchLabels() != nil {
		if o.MatchLabels == nil {
			o.MatchLabels = make(map[string]string, len(source.GetMatchLabels()))
		}
		for key, sourceValue := range source.GetMatchLabels() {
			o.MatchLabels[key] = sourceValue
		}
	}
	sourceMatchExpressions, origMatchExpressionsKeys := map[string]*LabelSelectorRequirement{}, map[string]bool{}
	newMatchExpressions := make([]*LabelSelectorRequirement, 0, len(o.MatchExpressions))
	for _, sourceValue := range source.GetMatchExpressions() {
		key := fmt.Sprintf("%s", sourceValue.GetKey())
		sourceMatchExpressions[key] = sourceValue
	}
	for _, origValue := range o.MatchExpressions {
		key := fmt.Sprintf("%s", origValue.GetKey())
		origMatchExpressionsKeys[key] = true
		sourceValue := sourceMatchExpressions[key]
		if sourceValue != nil {
			if origValue == nil {
				origValue = new(LabelSelectorRequirement)
			}
			origValue.Merge(sourceValue)
		}
		newMatchExpressions = append(newMatchExpressions, origValue)
	}
	for key, sourceValue := range sourceMatchExpressions {
		if origMatchExpressionsKeys[key] == false {
			newMatchExpressions = append(newMatchExpressions, sourceValue.Clone())
		}
	}
	o.MatchExpressions = newMatchExpressions

}

func (o *LabelSelector) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*LabelSelector))
}

func (o *LabelSelectorRequirement) GotenObjectExt() {}

func (o *LabelSelectorRequirement) MakeFullFieldMask() *LabelSelectorRequirement_FieldMask {
	return FullLabelSelectorRequirement_FieldMask()
}

func (o *LabelSelectorRequirement) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLabelSelectorRequirement_FieldMask()
}

func (o *LabelSelectorRequirement) MakeDiffFieldMask(other *LabelSelectorRequirement) *LabelSelectorRequirement_FieldMask {
	if o == nil && other == nil {
		return &LabelSelectorRequirement_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLabelSelectorRequirement_FieldMask()
	}

	res := &LabelSelectorRequirement_FieldMask{}
	if o.GetKey() != other.GetKey() {
		res.Paths = append(res.Paths, &LabelSelectorRequirement_FieldTerminalPath{selector: LabelSelectorRequirement_FieldPathSelectorKey})
	}
	if o.GetOperator() != other.GetOperator() {
		res.Paths = append(res.Paths, &LabelSelectorRequirement_FieldTerminalPath{selector: LabelSelectorRequirement_FieldPathSelectorOperator})
	}

	if len(o.GetValues()) == len(other.GetValues()) {
		for i, lValue := range o.GetValues() {
			rValue := other.GetValues()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &LabelSelectorRequirement_FieldTerminalPath{selector: LabelSelectorRequirement_FieldPathSelectorValues})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &LabelSelectorRequirement_FieldTerminalPath{selector: LabelSelectorRequirement_FieldPathSelectorValues})
	}
	return res
}

func (o *LabelSelectorRequirement) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*LabelSelectorRequirement))
}

func (o *LabelSelectorRequirement) Clone() *LabelSelectorRequirement {
	if o == nil {
		return nil
	}
	result := &LabelSelectorRequirement{}
	result.Key = o.Key
	result.Operator = o.Operator
	result.Values = make([]string, len(o.Values))
	for i, sourceValue := range o.Values {
		result.Values[i] = sourceValue
	}
	return result
}

func (o *LabelSelectorRequirement) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *LabelSelectorRequirement) Merge(source *LabelSelectorRequirement) {
	o.Key = source.GetKey()
	o.Operator = source.GetOperator()
	o.Values = make([]string, 0, len(source.GetValues()))
	for _, sourceValue := range source.GetValues() {
		var newDstElement string
		newDstElement = sourceValue
		o.Values = append(o.Values, newDstElement)
	}

}

func (o *LabelSelectorRequirement) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*LabelSelectorRequirement))
}

func (o *OwnerReference) GotenObjectExt() {}

func (o *OwnerReference) MakeFullFieldMask() *OwnerReference_FieldMask {
	return FullOwnerReference_FieldMask()
}

func (o *OwnerReference) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullOwnerReference_FieldMask()
}

func (o *OwnerReference) MakeDiffFieldMask(other *OwnerReference) *OwnerReference_FieldMask {
	if o == nil && other == nil {
		return &OwnerReference_FieldMask{}
	}
	if o == nil || other == nil {
		return FullOwnerReference_FieldMask()
	}

	res := &OwnerReference_FieldMask{}
	if o.GetKind() != other.GetKind() {
		res.Paths = append(res.Paths, &OwnerReference_FieldTerminalPath{selector: OwnerReference_FieldPathSelectorKind})
	}
	if o.GetVersion() != other.GetVersion() {
		res.Paths = append(res.Paths, &OwnerReference_FieldTerminalPath{selector: OwnerReference_FieldPathSelectorVersion})
	}
	if o.GetName() != other.GetName() {
		res.Paths = append(res.Paths, &OwnerReference_FieldTerminalPath{selector: OwnerReference_FieldPathSelectorName})
	}
	if o.GetRegion() != other.GetRegion() {
		res.Paths = append(res.Paths, &OwnerReference_FieldTerminalPath{selector: OwnerReference_FieldPathSelectorRegion})
	}
	if o.GetController() != other.GetController() {
		res.Paths = append(res.Paths, &OwnerReference_FieldTerminalPath{selector: OwnerReference_FieldPathSelectorController})
	}
	if o.GetRequiresOwnerReference() != other.GetRequiresOwnerReference() {
		res.Paths = append(res.Paths, &OwnerReference_FieldTerminalPath{selector: OwnerReference_FieldPathSelectorRequiresOwnerReference})
	}
	return res
}

func (o *OwnerReference) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*OwnerReference))
}

func (o *OwnerReference) Clone() *OwnerReference {
	if o == nil {
		return nil
	}
	result := &OwnerReference{}
	result.Kind = o.Kind
	result.Version = o.Version
	result.Name = o.Name
	result.Region = o.Region
	result.Controller = o.Controller
	result.RequiresOwnerReference = o.RequiresOwnerReference
	return result
}

func (o *OwnerReference) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *OwnerReference) Merge(source *OwnerReference) {
	o.Kind = source.GetKind()
	o.Version = source.GetVersion()
	o.Name = source.GetName()
	o.Region = source.GetRegion()
	o.Controller = source.GetController()
	o.RequiresOwnerReference = source.GetRequiresOwnerReference()
}

func (o *OwnerReference) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*OwnerReference))
}

func (o *SyncingMeta) GotenObjectExt() {}

func (o *SyncingMeta) MakeFullFieldMask() *SyncingMeta_FieldMask {
	return FullSyncingMeta_FieldMask()
}

func (o *SyncingMeta) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullSyncingMeta_FieldMask()
}

func (o *SyncingMeta) MakeDiffFieldMask(other *SyncingMeta) *SyncingMeta_FieldMask {
	if o == nil && other == nil {
		return &SyncingMeta_FieldMask{}
	}
	if o == nil || other == nil {
		return FullSyncingMeta_FieldMask()
	}

	res := &SyncingMeta_FieldMask{}
	if o.GetOwningRegion() != other.GetOwningRegion() {
		res.Paths = append(res.Paths, &SyncingMeta_FieldTerminalPath{selector: SyncingMeta_FieldPathSelectorOwningRegion})
	}

	if len(o.GetRegions()) == len(other.GetRegions()) {
		for i, lValue := range o.GetRegions() {
			rValue := other.GetRegions()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &SyncingMeta_FieldTerminalPath{selector: SyncingMeta_FieldPathSelectorRegions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &SyncingMeta_FieldTerminalPath{selector: SyncingMeta_FieldPathSelectorRegions})
	}
	return res
}

func (o *SyncingMeta) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*SyncingMeta))
}

func (o *SyncingMeta) Clone() *SyncingMeta {
	if o == nil {
		return nil
	}
	result := &SyncingMeta{}
	result.OwningRegion = o.OwningRegion
	result.Regions = make([]string, len(o.Regions))
	for i, sourceValue := range o.Regions {
		result.Regions[i] = sourceValue
	}
	return result
}

func (o *SyncingMeta) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *SyncingMeta) Merge(source *SyncingMeta) {
	o.OwningRegion = source.GetOwningRegion()
	for _, sourceValue := range source.GetRegions() {
		exists := false
		for _, currentValue := range o.Regions {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.Regions = append(o.Regions, newDstElement)
		}
	}

}

func (o *SyncingMeta) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*SyncingMeta))
}

func (o *Lifecycle) GotenObjectExt() {}

func (o *Lifecycle) MakeFullFieldMask() *Lifecycle_FieldMask {
	return FullLifecycle_FieldMask()
}

func (o *Lifecycle) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullLifecycle_FieldMask()
}

func (o *Lifecycle) MakeDiffFieldMask(other *Lifecycle) *Lifecycle_FieldMask {
	if o == nil && other == nil {
		return &Lifecycle_FieldMask{}
	}
	if o == nil || other == nil {
		return FullLifecycle_FieldMask()
	}

	res := &Lifecycle_FieldMask{}
	if o.GetState() != other.GetState() {
		res.Paths = append(res.Paths, &Lifecycle_FieldTerminalPath{selector: Lifecycle_FieldPathSelectorState})
	}
	if o.GetBlockDeletion() != other.GetBlockDeletion() {
		res.Paths = append(res.Paths, &Lifecycle_FieldTerminalPath{selector: Lifecycle_FieldPathSelectorBlockDeletion})
	}
	return res
}

func (o *Lifecycle) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Lifecycle))
}

func (o *Lifecycle) Clone() *Lifecycle {
	if o == nil {
		return nil
	}
	result := &Lifecycle{}
	result.State = o.State
	result.BlockDeletion = o.BlockDeletion
	return result
}

func (o *Lifecycle) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Lifecycle) Merge(source *Lifecycle) {
	o.State = source.GetState()
	o.BlockDeletion = source.GetBlockDeletion()
}

func (o *Lifecycle) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Lifecycle))
}

func (o *ServicesInfo) GotenObjectExt() {}

func (o *ServicesInfo) MakeFullFieldMask() *ServicesInfo_FieldMask {
	return FullServicesInfo_FieldMask()
}

func (o *ServicesInfo) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullServicesInfo_FieldMask()
}

func (o *ServicesInfo) MakeDiffFieldMask(other *ServicesInfo) *ServicesInfo_FieldMask {
	if o == nil && other == nil {
		return &ServicesInfo_FieldMask{}
	}
	if o == nil || other == nil {
		return FullServicesInfo_FieldMask()
	}

	res := &ServicesInfo_FieldMask{}
	if o.GetOwningService() != other.GetOwningService() {
		res.Paths = append(res.Paths, &ServicesInfo_FieldTerminalPath{selector: ServicesInfo_FieldPathSelectorOwningService})
	}

	if len(o.GetAllowedServices()) == len(other.GetAllowedServices()) {
		for i, lValue := range o.GetAllowedServices() {
			rValue := other.GetAllowedServices()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &ServicesInfo_FieldTerminalPath{selector: ServicesInfo_FieldPathSelectorAllowedServices})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &ServicesInfo_FieldTerminalPath{selector: ServicesInfo_FieldPathSelectorAllowedServices})
	}
	return res
}

func (o *ServicesInfo) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ServicesInfo))
}

func (o *ServicesInfo) Clone() *ServicesInfo {
	if o == nil {
		return nil
	}
	result := &ServicesInfo{}
	result.OwningService = o.OwningService
	result.AllowedServices = make([]string, len(o.AllowedServices))
	for i, sourceValue := range o.AllowedServices {
		result.AllowedServices[i] = sourceValue
	}
	return result
}

func (o *ServicesInfo) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ServicesInfo) Merge(source *ServicesInfo) {
	o.OwningService = source.GetOwningService()
	for _, sourceValue := range source.GetAllowedServices() {
		exists := false
		for _, currentValue := range o.AllowedServices {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.AllowedServices = append(o.AllowedServices, newDstElement)
		}
	}

}

func (o *ServicesInfo) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ServicesInfo))
}
