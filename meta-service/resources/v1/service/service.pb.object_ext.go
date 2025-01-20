// Code generated by protoc-gen-goten-object
// File: goten/meta-service/proto/v1/service.proto
// DO NOT EDIT!!!

package service

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	common "github.com/cloudwan/goten-sdk/meta-service/resources/v1/common"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	multi_region_policy "github.com/cloudwan/goten-sdk/types/multi_region_policy"
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
	_ = &common.LabelledDomain{}
	_ = &meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
)

func (o *Service) GotenObjectExt() {}

func (o *Service) MakeFullFieldMask() *Service_FieldMask {
	return FullService_FieldMask()
}

func (o *Service) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullService_FieldMask()
}

func (o *Service) MakeDiffFieldMask(other *Service) *Service_FieldMask {
	if o == nil && other == nil {
		return &Service_FieldMask{}
	}
	if o == nil || other == nil {
		return FullService_FieldMask()
	}

	res := &Service_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Service_FieldSubPath{selector: Service_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetMultiRegionPolicy().MakeDiffFieldMask(other.GetMultiRegionPolicy())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorMultiRegionPolicy})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Service_FieldSubPath{selector: Service_FieldPathSelectorMultiRegionPolicy, subPath: subpath})
			}
		}
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorDisplayName})
	}

	if len(o.GetAllVersions()) == len(other.GetAllVersions()) {
		for i, lValue := range o.GetAllVersions() {
			rValue := other.GetAllVersions()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorAllVersions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorAllVersions})
	}
	if o.GetGlobalDomain() != other.GetGlobalDomain() {
		res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorGlobalDomain})
	}

	if len(o.GetLabelledDomains()) == len(other.GetLabelledDomains()) {
		for i, lValue := range o.GetLabelledDomains() {
			rValue := other.GetLabelledDomains()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorLabelledDomains})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorLabelledDomains})
	}
	if o.GetLeadingService().String() != other.GetLeadingService().String() {
		res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorLeadingService})
	}

	if len(o.GetImportedServices()) == len(other.GetImportedServices()) {
		for i, lValue := range o.GetImportedServices() {
			rValue := other.GetImportedServices()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorImportedServices})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorImportedServices})
	}

	if len(o.GetUsedServices()) == len(other.GetUsedServices()) {
		for i, lValue := range o.GetUsedServices() {
			rValue := other.GetUsedServices()[i]
			if lValue.String() != rValue.String() {
				res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorUsedServices})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorUsedServices})
	}

	if len(o.GetImportedVersions()) == len(other.GetImportedVersions()) {
		for i, lValue := range o.GetImportedVersions() {
			rValue := other.GetImportedVersions()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorImportedVersions})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorImportedVersions})
	}
	if o.GetEnvRegistryGeneration() != other.GetEnvRegistryGeneration() {
		res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorEnvRegistryGeneration})
	}
	if o.GetAutomaticVersionSwitch() != other.GetAutomaticVersionSwitch() {
		res.Paths = append(res.Paths, &Service_FieldTerminalPath{selector: Service_FieldPathSelectorAutomaticVersionSwitch})
	}
	return res
}

func (o *Service) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Service))
}

func (o *Service) Clone() *Service {
	if o == nil {
		return nil
	}
	result := &Service{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &Name{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Metadata = o.Metadata.Clone()
	result.MultiRegionPolicy = o.MultiRegionPolicy.Clone()
	result.DisplayName = o.DisplayName
	result.AllVersions = make([]string, len(o.AllVersions))
	for i, sourceValue := range o.AllVersions {
		result.AllVersions[i] = sourceValue
	}
	result.GlobalDomain = o.GlobalDomain
	result.LabelledDomains = make([]*common.LabelledDomain, len(o.LabelledDomains))
	for i, sourceValue := range o.LabelledDomains {
		result.LabelledDomains[i] = sourceValue.Clone()
	}
	if o.LeadingService == nil {
		result.LeadingService = nil
	} else if data, err := o.LeadingService.ProtoString(); err != nil {
		panic(err)
	} else {
		result.LeadingService = &Name{}
		if err := result.LeadingService.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.ImportedServices = make([]*Reference, len(o.ImportedServices))
	for i, sourceValue := range o.ImportedServices {
		if sourceValue == nil {
			result.ImportedServices[i] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.ImportedServices[i] = &Reference{}
			if err := result.ImportedServices[i].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
	result.UsedServices = make([]*Reference, len(o.UsedServices))
	for i, sourceValue := range o.UsedServices {
		if sourceValue == nil {
			result.UsedServices[i] = nil
		} else if data, err := sourceValue.ProtoString(); err != nil {
			panic(err)
		} else {
			result.UsedServices[i] = &Reference{}
			if err := result.UsedServices[i].ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	}
	result.ImportedVersions = make([]*Service_ImportedVersions, len(o.ImportedVersions))
	for i, sourceValue := range o.ImportedVersions {
		result.ImportedVersions[i] = sourceValue.Clone()
	}
	result.EnvRegistryGeneration = o.EnvRegistryGeneration
	result.AutomaticVersionSwitch = o.AutomaticVersionSwitch
	return result
}

func (o *Service) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Service) Merge(source *Service) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &Name{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
	if source.GetMultiRegionPolicy() != nil {
		if o.MultiRegionPolicy == nil {
			o.MultiRegionPolicy = new(multi_region_policy.MultiRegionPolicy)
		}
		o.MultiRegionPolicy.Merge(source.GetMultiRegionPolicy())
	}
	o.DisplayName = source.GetDisplayName()
	for _, sourceValue := range source.GetAllVersions() {
		exists := false
		for _, currentValue := range o.AllVersions {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.AllVersions = append(o.AllVersions, newDstElement)
		}
	}

	o.GlobalDomain = source.GetGlobalDomain()
	for _, sourceValue := range source.GetLabelledDomains() {
		exists := false
		for _, currentValue := range o.LabelledDomains {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *common.LabelledDomain
			if sourceValue != nil {
				newDstElement = new(common.LabelledDomain)
				newDstElement.Merge(sourceValue)
			}
			o.LabelledDomains = append(o.LabelledDomains, newDstElement)
		}
	}

	if source.GetLeadingService() != nil {
		if data, err := source.GetLeadingService().ProtoString(); err != nil {
			panic(err)
		} else {
			o.LeadingService = &Name{}
			if err := o.LeadingService.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.LeadingService = nil
	}
	for _, sourceValue := range source.GetImportedServices() {
		exists := false
		for _, currentValue := range o.ImportedServices {
			leftProtoStr, _ := currentValue.ProtoString()
			rightProtoStr, _ := sourceValue.ProtoString()
			if leftProtoStr == rightProtoStr {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *Reference
			if sourceValue != nil {
				if data, err := sourceValue.ProtoString(); err != nil {
					panic(err)
				} else {
					newDstElement = &Reference{}
					if err := newDstElement.ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			}
			o.ImportedServices = append(o.ImportedServices, newDstElement)
		}
	}

	for _, sourceValue := range source.GetUsedServices() {
		exists := false
		for _, currentValue := range o.UsedServices {
			leftProtoStr, _ := currentValue.ProtoString()
			rightProtoStr, _ := sourceValue.ProtoString()
			if leftProtoStr == rightProtoStr {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *Reference
			if sourceValue != nil {
				if data, err := sourceValue.ProtoString(); err != nil {
					panic(err)
				} else {
					newDstElement = &Reference{}
					if err := newDstElement.ParseProtoString(data); err != nil {
						panic(err)
					}
				}
			}
			o.UsedServices = append(o.UsedServices, newDstElement)
		}
	}

	for _, sourceValue := range source.GetImportedVersions() {
		exists := false
		for _, currentValue := range o.ImportedVersions {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *Service_ImportedVersions
			if sourceValue != nil {
				newDstElement = new(Service_ImportedVersions)
				newDstElement.Merge(sourceValue)
			}
			o.ImportedVersions = append(o.ImportedVersions, newDstElement)
		}
	}

	o.EnvRegistryGeneration = source.GetEnvRegistryGeneration()
	o.AutomaticVersionSwitch = source.GetAutomaticVersionSwitch()
}

func (o *Service) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Service))
}

func (o *Service_ImportedVersions) GotenObjectExt() {}

func (o *Service_ImportedVersions) MakeFullFieldMask() *Service_ImportedVersions_FieldMask {
	return FullService_ImportedVersions_FieldMask()
}

func (o *Service_ImportedVersions) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullService_ImportedVersions_FieldMask()
}

func (o *Service_ImportedVersions) MakeDiffFieldMask(other *Service_ImportedVersions) *Service_ImportedVersions_FieldMask {
	if o == nil && other == nil {
		return &Service_ImportedVersions_FieldMask{}
	}
	if o == nil || other == nil {
		return FullService_ImportedVersions_FieldMask()
	}

	res := &Service_ImportedVersions_FieldMask{}
	if o.GetTargetService().String() != other.GetTargetService().String() {
		res.Paths = append(res.Paths, &ServiceImportedVersions_FieldTerminalPath{selector: ServiceImportedVersions_FieldPathSelectorTargetService})
	}
	if o.GetTargetServiceVersion() != other.GetTargetServiceVersion() {
		res.Paths = append(res.Paths, &ServiceImportedVersions_FieldTerminalPath{selector: ServiceImportedVersions_FieldPathSelectorTargetServiceVersion})
	}
	if o.GetCurrentServiceVersion() != other.GetCurrentServiceVersion() {
		res.Paths = append(res.Paths, &ServiceImportedVersions_FieldTerminalPath{selector: ServiceImportedVersions_FieldPathSelectorCurrentServiceVersion})
	}
	return res
}

func (o *Service_ImportedVersions) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Service_ImportedVersions))
}

func (o *Service_ImportedVersions) Clone() *Service_ImportedVersions {
	if o == nil {
		return nil
	}
	result := &Service_ImportedVersions{}
	if o.TargetService == nil {
		result.TargetService = nil
	} else if data, err := o.TargetService.ProtoString(); err != nil {
		panic(err)
	} else {
		result.TargetService = &Reference{}
		if err := result.TargetService.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.TargetServiceVersion = o.TargetServiceVersion
	result.CurrentServiceVersion = o.CurrentServiceVersion
	return result
}

func (o *Service_ImportedVersions) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Service_ImportedVersions) Merge(source *Service_ImportedVersions) {
	if source.GetTargetService() != nil {
		if data, err := source.GetTargetService().ProtoString(); err != nil {
			panic(err)
		} else {
			o.TargetService = &Reference{}
			if err := o.TargetService.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.TargetService = nil
	}
	o.TargetServiceVersion = source.GetTargetServiceVersion()
	o.CurrentServiceVersion = source.GetCurrentServiceVersion()
}

func (o *Service_ImportedVersions) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Service_ImportedVersions))
}
