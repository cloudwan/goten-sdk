// Code generated by protoc-gen-goten-resource
// Resource: Deployment
// DO NOT EDIT!!!

package deployment

import (
	googlefieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/cloudwan/goten-sdk/types/view"
)

// proto imports
import (
	region "github.com/cloudwan/goten-sdk/meta-service/resources/v1/region"
	service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
	meta "github.com/cloudwan/goten-sdk/types/meta"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// ensure the imports are used
var (
	_ = googlefieldmaskpb.FieldMask{}

	_ = view.View_UNSPECIFIED
)

// make sure we're using proto imports
var (
	_ = &structpb.Struct{}
	_ = &region.Region{}
	_ = &service.Service{}
	_ = &meta.Meta{}
)

func ResourceViewFieldMask(viewName view.View, extraMask *Deployment_FieldMask) *Deployment_FieldMask {
	protoFieldMask := &googlefieldmaskpb.FieldMask{}

	switch viewName {
	case view.View_UNSPECIFIED:
		return extraMask
	case view.View_FULL:
		return nil
	case view.View_NAME:
		protoFieldMask.Paths = append(protoFieldMask.Paths, "name")
		break
	default:
		return extraMask
	}
	if extraMask != nil {
		protoFieldMask.Paths = append(protoFieldMask.Paths, extraMask.ToProtoFieldMask().Paths...)
	}
	res := &Deployment_FieldMask{}
	_ = res.FromProtoFieldMask(protoFieldMask)
	return res
}
