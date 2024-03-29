// Code generated by protoc-gen-goten-access
// Service: Meta
// DO NOT EDIT!!!

package meta_access

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	deployment_access "github.com/cloudwan/goten-sdk/meta-service/access/v1/deployment"
	region_access "github.com/cloudwan/goten-sdk/meta-service/access/v1/region"
	resource_access "github.com/cloudwan/goten-sdk/meta-service/access/v1/resource"
	service_access "github.com/cloudwan/goten-sdk/meta-service/access/v1/service"
	meta_client "github.com/cloudwan/goten-sdk/meta-service/client/v1/meta"
	deployment "github.com/cloudwan/goten-sdk/meta-service/resources/v1/deployment"
	region "github.com/cloudwan/goten-sdk/meta-service/resources/v1/region"
	resource "github.com/cloudwan/goten-sdk/meta-service/resources/v1/resource"
	service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
)

type MetaApiAccess interface {
	gotenresource.Access

	deployment.DeploymentAccess
	region.RegionAccess
	resource.ResourceAccess
	service.ServiceAccess
}

type apiMetaAccess struct {
	gotenresource.Access

	deployment.DeploymentAccess
	region.RegionAccess
	resource.ResourceAccess
	service.ServiceAccess
}

func NewApiAccess(client meta_client.MetaClient) MetaApiAccess {

	deploymentAccess := deployment_access.NewApiDeploymentAccess(client)
	regionAccess := region_access.NewApiRegionAccess(client)
	resourceAccess := resource_access.NewApiResourceAccess(client)
	serviceAccess := service_access.NewApiServiceAccess(client)

	return &apiMetaAccess{
		Access: gotenresource.NewCompositeAccess(

			deployment.AsAnyCastAccess(deploymentAccess),
			region.AsAnyCastAccess(regionAccess),
			resource.AsAnyCastAccess(resourceAccess),
			service.AsAnyCastAccess(serviceAccess),
		),

		DeploymentAccess: deploymentAccess,
		RegionAccess:     regionAccess,
		ResourceAccess:   resourceAccess,
		ServiceAccess:    serviceAccess,
	}
}
