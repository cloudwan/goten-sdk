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
	common "github.com/cloudwan/goten-sdk/meta-service/resources/v1/common"
	region "github.com/cloudwan/goten-sdk/meta-service/resources/v1/region"
	service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	structpb "google.golang.org/protobuf/types/known/structpb"
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
	_ = &structpb.Struct{}
	_ = &common.LabelledDomain{}
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

	if len(o.GetLabelledDomains()) == len(other.GetLabelledDomains()) {
		for i, lValue := range o.GetLabelledDomains() {
			rValue := other.GetLabelledDomains()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorLabelledDomains})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorLabelledDomains})
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
	if o.GetDbDataVersion() != other.GetDbDataVersion() {
		res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorDbDataVersion})
	}
	if o.GetDbLocationTag() != other.GetDbLocationTag() {
		res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorDbLocationTag})
	}
	if o.GetAutomaticVersionSwitch() != other.GetAutomaticVersionSwitch() {
		res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorAutomaticVersionSwitch})
	}
	{
		subMask := o.GetAvailableUpgrade().MakeDiffFieldMask(other.GetAvailableUpgrade())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorAvailableUpgrade})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &Deployment_FieldSubPath{selector: Deployment_FieldPathSelectorAvailableUpgrade, subPath: subpath})
			}
		}
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

	if len(o.GetDataUpdateStatuses()) == len(other.GetDataUpdateStatuses()) {
		for i, lValue := range o.GetDataUpdateStatuses() {
			rValue := other.GetDataUpdateStatuses()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorDataUpdateStatuses})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &Deployment_FieldTerminalPath{selector: Deployment_FieldPathSelectorDataUpdateStatuses})
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
	result.LabelledDomains = make([]*common.LabelledDomain, len(o.LabelledDomains))
	for i, sourceValue := range o.LabelledDomains {
		result.LabelledDomains[i] = sourceValue.Clone()
	}
	result.LocalNetworkId = o.LocalNetworkId
	result.Location = o.Location.Clone()
	result.IsDisabled = o.IsDisabled
	result.EnvRegistryGeneration = o.EnvRegistryGeneration
	result.CurrentVersion = o.CurrentVersion
	result.DbDataVersion = o.DbDataVersion
	result.DbLocationTag = o.DbLocationTag
	result.AutomaticVersionSwitch = o.AutomaticVersionSwitch
	result.AvailableUpgrade = o.AvailableUpgrade.Clone()
	result.UpgradeState = o.UpgradeState.Clone()
	result.DataUpdateStatuses = make([]*Deployment_DbUpdateTaskStatus, len(o.DataUpdateStatuses))
	for i, sourceValue := range o.DataUpdateStatuses {
		result.DataUpdateStatuses[i] = sourceValue.Clone()
	}
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
	o.DbDataVersion = source.GetDbDataVersion()
	o.DbLocationTag = source.GetDbLocationTag()
	o.AutomaticVersionSwitch = source.GetAutomaticVersionSwitch()
	if source.GetAvailableUpgrade() != nil {
		if o.AvailableUpgrade == nil {
			o.AvailableUpgrade = new(Deployment_AvailableUpgrade)
		}
		o.AvailableUpgrade.Merge(source.GetAvailableUpgrade())
	}
	if source.GetUpgradeState() != nil {
		if o.UpgradeState == nil {
			o.UpgradeState = new(Deployment_UpgradeState)
		}
		o.UpgradeState.Merge(source.GetUpgradeState())
	}
	for _, sourceValue := range source.GetDataUpdateStatuses() {
		exists := false
		for _, currentValue := range o.DataUpdateStatuses {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *Deployment_DbUpdateTaskStatus
			if sourceValue != nil {
				newDstElement = new(Deployment_DbUpdateTaskStatus)
				newDstElement.Merge(sourceValue)
			}
			o.DataUpdateStatuses = append(o.DataUpdateStatuses, newDstElement)
		}
	}

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
	if o.GetDbDataTargetVersion() != other.GetDbDataTargetVersion() {
		res.Paths = append(res.Paths, &DeploymentUpgradeState_FieldTerminalPath{selector: DeploymentUpgradeState_FieldPathSelectorDbDataTargetVersion})
	}
	if o.GetDbTargetLocationTag() != other.GetDbTargetLocationTag() {
		res.Paths = append(res.Paths, &DeploymentUpgradeState_FieldTerminalPath{selector: DeploymentUpgradeState_FieldPathSelectorDbTargetLocationTag})
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
	result.DbDataTargetVersion = o.DbDataTargetVersion
	result.DbTargetLocationTag = o.DbTargetLocationTag
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
	o.DbDataTargetVersion = source.GetDbDataTargetVersion()
	o.DbTargetLocationTag = source.GetDbTargetLocationTag()
}

func (o *Deployment_UpgradeState) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Deployment_UpgradeState))
}

func (o *Deployment_AvailableUpgrade) GotenObjectExt() {}

func (o *Deployment_AvailableUpgrade) MakeFullFieldMask() *Deployment_AvailableUpgrade_FieldMask {
	return FullDeployment_AvailableUpgrade_FieldMask()
}

func (o *Deployment_AvailableUpgrade) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullDeployment_AvailableUpgrade_FieldMask()
}

func (o *Deployment_AvailableUpgrade) MakeDiffFieldMask(other *Deployment_AvailableUpgrade) *Deployment_AvailableUpgrade_FieldMask {
	if o == nil && other == nil {
		return &Deployment_AvailableUpgrade_FieldMask{}
	}
	if o == nil || other == nil {
		return FullDeployment_AvailableUpgrade_FieldMask()
	}

	res := &Deployment_AvailableUpgrade_FieldMask{}
	if o.GetApiVersion() != other.GetApiVersion() {
		res.Paths = append(res.Paths, &DeploymentAvailableUpgrade_FieldTerminalPath{selector: DeploymentAvailableUpgrade_FieldPathSelectorApiVersion})
	}
	if o.GetDbDataVersion() != other.GetDbDataVersion() {
		res.Paths = append(res.Paths, &DeploymentAvailableUpgrade_FieldTerminalPath{selector: DeploymentAvailableUpgrade_FieldPathSelectorDbDataVersion})
	}
	if o.GetDbLocationTag() != other.GetDbLocationTag() {
		res.Paths = append(res.Paths, &DeploymentAvailableUpgrade_FieldTerminalPath{selector: DeploymentAvailableUpgrade_FieldPathSelectorDbLocationTag})
	}
	if o.GetTotalShardsCount() != other.GetTotalShardsCount() {
		res.Paths = append(res.Paths, &DeploymentAvailableUpgrade_FieldTerminalPath{selector: DeploymentAvailableUpgrade_FieldPathSelectorTotalShardsCount})
	}
	return res
}

func (o *Deployment_AvailableUpgrade) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Deployment_AvailableUpgrade))
}

func (o *Deployment_AvailableUpgrade) Clone() *Deployment_AvailableUpgrade {
	if o == nil {
		return nil
	}
	result := &Deployment_AvailableUpgrade{}
	result.ApiVersion = o.ApiVersion
	result.DbDataVersion = o.DbDataVersion
	result.DbLocationTag = o.DbLocationTag
	result.TotalShardsCount = o.TotalShardsCount
	return result
}

func (o *Deployment_AvailableUpgrade) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Deployment_AvailableUpgrade) Merge(source *Deployment_AvailableUpgrade) {
	o.ApiVersion = source.GetApiVersion()
	o.DbDataVersion = source.GetDbDataVersion()
	o.DbLocationTag = source.GetDbLocationTag()
	o.TotalShardsCount = source.GetTotalShardsCount()
}

func (o *Deployment_AvailableUpgrade) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Deployment_AvailableUpgrade))
}

func (o *Deployment_DbUpdateTaskStatus) GotenObjectExt() {}

func (o *Deployment_DbUpdateTaskStatus) MakeFullFieldMask() *Deployment_DbUpdateTaskStatus_FieldMask {
	return FullDeployment_DbUpdateTaskStatus_FieldMask()
}

func (o *Deployment_DbUpdateTaskStatus) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullDeployment_DbUpdateTaskStatus_FieldMask()
}

func (o *Deployment_DbUpdateTaskStatus) MakeDiffFieldMask(other *Deployment_DbUpdateTaskStatus) *Deployment_DbUpdateTaskStatus_FieldMask {
	if o == nil && other == nil {
		return &Deployment_DbUpdateTaskStatus_FieldMask{}
	}
	if o == nil || other == nil {
		return FullDeployment_DbUpdateTaskStatus_FieldMask()
	}

	res := &Deployment_DbUpdateTaskStatus_FieldMask{}
	if o.GetTaskTag() != other.GetTaskTag() {
		res.Paths = append(res.Paths, &DeploymentDbUpdateTaskStatus_FieldTerminalPath{selector: DeploymentDbUpdateTaskStatus_FieldPathSelectorTaskTag})
	}
	if o.GetShardsCount() != other.GetShardsCount() {
		res.Paths = append(res.Paths, &DeploymentDbUpdateTaskStatus_FieldTerminalPath{selector: DeploymentDbUpdateTaskStatus_FieldPathSelectorShardsCount})
	}

	if len(o.GetReadyShards()) == len(other.GetReadyShards()) {
		for i, lValue := range o.GetReadyShards() {
			rValue := other.GetReadyShards()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &DeploymentDbUpdateTaskStatus_FieldTerminalPath{selector: DeploymentDbUpdateTaskStatus_FieldPathSelectorReadyShards})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &DeploymentDbUpdateTaskStatus_FieldTerminalPath{selector: DeploymentDbUpdateTaskStatus_FieldPathSelectorReadyShards})
	}
	if o.GetTargetVersion() != other.GetTargetVersion() {
		res.Paths = append(res.Paths, &DeploymentDbUpdateTaskStatus_FieldTerminalPath{selector: DeploymentDbUpdateTaskStatus_FieldPathSelectorTargetVersion})
	}

	if len(o.GetProgressBar()) == len(other.GetProgressBar()) {
		for i, lValue := range o.GetProgressBar() {
			rValue := other.GetProgressBar()[i]
			if !proto.Equal(lValue, rValue) {
				res.Paths = append(res.Paths, &DeploymentDbUpdateTaskStatus_FieldTerminalPath{selector: DeploymentDbUpdateTaskStatus_FieldPathSelectorProgressBar})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &DeploymentDbUpdateTaskStatus_FieldTerminalPath{selector: DeploymentDbUpdateTaskStatus_FieldPathSelectorProgressBar})
	}
	return res
}

func (o *Deployment_DbUpdateTaskStatus) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*Deployment_DbUpdateTaskStatus))
}

func (o *Deployment_DbUpdateTaskStatus) Clone() *Deployment_DbUpdateTaskStatus {
	if o == nil {
		return nil
	}
	result := &Deployment_DbUpdateTaskStatus{}
	result.TaskTag = o.TaskTag
	result.ShardsCount = o.ShardsCount
	result.ReadyShards = make([]int64, len(o.ReadyShards))
	for i, sourceValue := range o.ReadyShards {
		result.ReadyShards[i] = sourceValue
	}
	result.TargetVersion = o.TargetVersion
	result.ProgressBar = make([]*structpb.Struct, len(o.ProgressBar))
	for i, sourceValue := range o.ProgressBar {
		result.ProgressBar[i] = proto.Clone(sourceValue).(*structpb.Struct)
	}
	return result
}

func (o *Deployment_DbUpdateTaskStatus) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *Deployment_DbUpdateTaskStatus) Merge(source *Deployment_DbUpdateTaskStatus) {
	o.TaskTag = source.GetTaskTag()
	o.ShardsCount = source.GetShardsCount()
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

	o.TargetVersion = source.GetTargetVersion()
	for _, sourceValue := range source.GetProgressBar() {
		exists := false
		for _, currentValue := range o.ProgressBar {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *structpb.Struct
			if sourceValue != nil {
				newDstElement = new(structpb.Struct)
				proto.Merge(newDstElement, sourceValue)
			}
			o.ProgressBar = append(o.ProgressBar, newDstElement)
		}
	}

}

func (o *Deployment_DbUpdateTaskStatus) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*Deployment_DbUpdateTaskStatus))
}
