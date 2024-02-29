package resource

import (
	"fmt"
	"strings"

	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

func MakeNameFromFullyQualifiedName(fqn string) *Name {
	elements := strings.Split(fqn, "/")
	return NewNameBuilder().SetServiceId(elements[0]).SetId(elements[1]).Name()
}

func MakeNameFromResourceDescriptor(descriptor gotenresource.Descriptor) *Name {
	return NewNameBuilder().
		SetServiceId(descriptor.GetResourceTypeName().ServiceDomain()).
		SetId(descriptor.GetResourceTypeName().Singular()).Name()
}

func (n *Name) FormResourceFullyQualifiedTypeName() string {
	return fmt.Sprintf("%s/%s", n.ServiceId, n.ResourceId)
}
