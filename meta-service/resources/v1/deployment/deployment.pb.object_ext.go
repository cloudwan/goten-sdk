// Code generated by protoc-gen-goten-object
// File: goten/meta-service/proto/v1/deployment.proto
// DO NOT EDIT!!!

package deployment

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	region "github.com/cloudwan/goten-sdk/meta-service/resources/v1/region"
	service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
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
	_ = &region.Region{}
	_ = &service.Service{}
	_ = &meta.Meta{}
)

func (o *Deployment) GotenObjectExt() {}

func (o *Deployment) MakeFullFieldMask() *Deployment_FieldMask {
	return FullDeployment_FieldMask()
}

func (o *Deployment) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullDeployment_FieldMask()
}

func (o *Deployment) MakeDiffFieldMask(other *Deployment) *Deployment_FieldMask {
	if o == nil && other == nil {
		return &Deployment_FieldMask{}
	}
	if o == nil || other == nil {
		return FullDeployment_FieldMask()
	}

	res := &Deployment_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Deployment_FieldSubPath{selector: Deployment_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetServiceName() != other.GetServiceName() {
		res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorServiceName})
	}
	if o.GetRegion().String() != other.GetRegion().String() {
		res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorRegion})
	}
	if o.GetPublicDomain() != other.GetPublicDomain() {
		res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorPublicDomain})
	}
	if o.GetPrivateDomain() != other.GetPrivateDomain() {
		res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorPrivateDomain})
	}
	if o.GetLocalNetworkId() != other.GetLocalNetworkId() {
		res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorLocalNetworkId})
	}
	{
		subMask := o.GetLocation().MakeDiffFieldMask(other.GetLocation())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorLocation})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Deployment_FieldSubPath{selector: Deployment_FieldPathSelectorLocation, subPath: subpath})
			}
		}
	}
	if o.GetIsDisabled() != other.GetIsDisabled() {
		res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorIsDisabled})
	}
	if o.GetEnvRegistryGeneration() != other.GetEnvRegistryGeneration() {
		res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorEnvRegistryGeneration})
	}
	if o.GetCurrentVersion() != other.GetCurrentVersion() {
		res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorCurrentVersion})
	}
	if o.GetAutomaticVersionSwitch() != other.GetAutomaticVersionSwitch() {
		res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorAutomaticVersionSwitch})
	}
	{
		subMask := o.GetUpgradeState().MakeDiffFieldMask(other.GetUpgradeState())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorUpgradeState})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Deployment_FieldSubPath{selector: Deployment_FieldPathSelectorUpgradeState, subPath: subpath})
			}
		}
	}
	if o.GetAllowedServicesGeneration() != other.GetAllowedServicesGeneration() {
		res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorAllowedServicesGeneration})
	}
	return res
}

func (o *Deployment) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Deployment))
}

func (o *Deployment) Clone() *Deployment {
	if o == nil {
		return nil
	}
	result := &Deployment{}
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
	result.ServiceName = o.ServiceName
	if o.Region == nil {
		result.Region = nil
	} else if data, err := o.Region.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Region = &region.Reference{}
		if err := result.Region.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.PublicDomain = o.PublicDomain
	result.PrivateDomain = o.PrivateDomain
	result.LocalNetworkId = o.LocalNetworkId
	result.Location = o.Location.Clone()
	result.IsDisabled = o.IsDisabled
	result.EnvRegistryGeneration = o.EnvRegistryGeneration
	result.CurrentVersion = o.CurrentVersion
	result.AutomaticVersionSwitch = o.AutomaticVersionSwitch
	result.UpgradeState = o.UpgradeState.Clone()
	result.AllowedServicesGeneration = o.AllowedServicesGeneration
	return result
}

func (o *Deployment) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Deployment) Merge(source *Deployment) {
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
	o.ServiceName = source.GetServiceName()
	if source.GetRegion() != nil {
		if data, err := source.GetRegion().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Region = &region.Reference{}
			if err := o.Region.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Region = nil
	}
	o.PublicDomain = source.GetPublicDomain()
	o.PrivateDomain = source.GetPrivateDomain()
	o.LocalNetworkId = source.GetLocalNetworkId()
	if source.GetLocation() != nil {
		if o.Location == nil {
			o.Location = new(Deployment_Location)
		}
		o.Location.Merge(source.GetLocation())
	}
	o.IsDisabled = source.GetIsDisabled()
	o.EnvRegistryGeneration = source.GetEnvRegistryGeneration()
	o.CurrentVersion = source.GetCurrentVersion()
	o.AutomaticVersionSwitch = source.GetAutomaticVersionSwitch()
	if source.GetUpgradeState() != nil {
		if o.UpgradeState == nil {
			o.UpgradeState = new(Deployment_UpgradeState)
		}
		o.UpgradeState.Merge(source.GetUpgradeState())
	}
	o.AllowedServicesGeneration = source.GetAllowedServicesGeneration()
}

func (o *Deployment) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Deployment))
}

func (o *Deployment_Location) GotenObjectExt() {}

func (o *Deployment_Location) MakeFullFieldMask() *Deployment_Location_FieldMask {
	return FullDeployment_Location_FieldMask()
}

func (o *Deployment_Location) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullDeployment_Location_FieldMask()
}

func (o *Deployment_Location) MakeDiffFieldMask(other *Deployment_Location) *Deployment_Location_FieldMask {
	if o == nil && other == nil {
		return &Deployment_Location_FieldMask{}
	}
	if o == nil || other == nil {
		return FullDeployment_Location_FieldMask()
	}

	res := &Deployment_Location_FieldMask{}
	if o.GetContinent() != other.GetContinent() {
		res.Paths = append(res.Paths, &DeploymentLocation_FieldTerminalPath{selector: DeploymentLocation_FieldPathSelectorContinent})
	}
	if o.GetCountry() != other.GetCountry() {
		res.Paths = append(res.Paths, &DeploymentLocation_FieldTerminalPath{selector: DeploymentLocation_FieldPathSelectorCountry})
	}
	if o.GetAgglomeration() != other.GetAgglomeration() {
		res.Paths = append(res.Paths, &DeploymentLocation_FieldTerminalPath{selector: DeploymentLocation_FieldPathSelectorAgglomeration})
	}
	if o.GetCity() != other.GetCity() {
		res.Paths = append(res.Paths, &DeploymentLocation_FieldTerminalPath{selector: DeploymentLocation_FieldPathSelectorCity})
	}
	if o.GetCloud() != other.GetCloud() {
		res.Paths = append(res.Paths, &DeploymentLocation_FieldTerminalPath{selector: DeploymentLocation_FieldPathSelectorCloud})
	}
	return res
}

func (o *Deployment_Location) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Deployment_Location))
}

func (o *Deployment_Location) Clone() *Deployment_Location {
	if o == nil {
		return nil
	}
	result := &Deployment_Location{}
	result.Continent = o.Continent
	result.Country = o.Country
	result.Agglomeration = o.Agglomeration
	result.City = o.City
	result.Cloud = o.Cloud
	return result
}

func (o *Deployment_Location) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Deployment_Location) Merge(source *Deployment_Location) {
	o.Continent = source.GetContinent()
	o.Country = source.GetCountry()
	o.Agglomeration = source.GetAgglomeration()
	o.City = source.GetCity()
	o.Cloud = source.GetCloud()
}

func (o *Deployment_Location) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Deployment_Location))
}

func (o *Deployment_UpgradeState) GotenObjectExt() {}

func (o *Deployment_UpgradeState) MakeFullFieldMask() *Deployment_UpgradeState_FieldMask {
	return FullDeployment_UpgradeState_FieldMask()
}

func (o *Deployment_UpgradeState) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullDeployment_UpgradeState_FieldMask()
}

func (o *Deployment_UpgradeState) MakeDiffFieldMask(other *Deployment_UpgradeState) *Deployment_UpgradeState_FieldMask {
	if o == nil && other == nil {
		return &Deployment_UpgradeState_FieldMask{}
	}
	if o == nil || other == nil {
		return FullDeployment_UpgradeState_FieldMask()
	}

	res := &Deployment_UpgradeState_FieldMask{}
	if o.GetTargetVersion() != other.GetTargetVersion() {
		res.Paths = append(res.Paths, &DeploymentUpgradeState_FieldTerminalPath{selector: DeploymentUpgradeState_FieldPathSelectorTargetVersion})
	}

	if len(o.GetReadyShards()) == len(other.GetReadyShards()) {
		for i, lValue := range o.GetReadyShards() {
			rValue := other.GetReadyShards()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &DeploymentUpgradeState_FieldTerminalPath{selector: DeploymentUpgradeState_FieldPathSelectorReadyShards})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &DeploymentUpgradeState_FieldTerminalPath{selector: DeploymentUpgradeState_FieldPathSelectorReadyShards})
	}

	if len(o.GetPendingShards()) == len(other.GetPendingShards()) {
		for i, lValue := range o.GetPendingShards() {
			rValue := other.GetPendingShards()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &DeploymentUpgradeState_FieldTerminalPath{selector: DeploymentUpgradeState_FieldPathSelectorPendingShards})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &DeploymentUpgradeState_FieldTerminalPath{selector: DeploymentUpgradeState_FieldPathSelectorPendingShards})
	}
	if o.GetStage() != other.GetStage() {
		res.Paths = append(res.Paths, &DeploymentUpgradeState_FieldTerminalPath{selector: DeploymentUpgradeState_FieldPathSelectorStage})
	}
	return res
}

func (o *Deployment_UpgradeState) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Deployment_UpgradeState))
}

func (o *Deployment_UpgradeState) Clone() *Deployment_UpgradeState {
	if o == nil {
		return nil
	}
	result := &Deployment_UpgradeState{}
	result.TargetVersion = o.TargetVersion
	result.ReadyShards = make([]int64, len(o.ReadyShards))
	for i, sourceValue := range o.ReadyShards {
		result.ReadyShards[i] = sourceValue
	}
	result.PendingShards = make([]int64, len(o.PendingShards))
	for i, sourceValue := range o.PendingShards {
		result.PendingShards[i] = sourceValue
	}
	result.Stage = o.Stage
	return result
}

func (o *Deployment_UpgradeState) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Deployment_UpgradeState) Merge(source *Deployment_UpgradeState) {
	o.TargetVersion = source.GetTargetVersion()
	for _, sourceValue := range source.GetReadyShards() {
		exists := false
		for _, currentValue := range o.ReadyShards {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement int64
			newDstElement = sourceValue
			o.ReadyShards = append(o.ReadyShards, newDstElement)
		}
	}

	for _, sourceValue := range source.GetPendingShards() {
		exists := false
		for _, currentValue := range o.PendingShards {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement int64
			newDstElement = sourceValue
			o.PendingShards = append(o.PendingShards, newDstElement)
		}
	}

	o.Stage = source.GetStage()
}

func (o *Deployment_UpgradeState) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Deployment_UpgradeState))
}