// Code generated by protoc-gen-goten-resource
// Resource change: ServiceChange
// DO NOT EDIT!!!

package service

import (
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ensure the imports are used
var (
	_ = new(gotenresource.ListQuery)
)

// make sure we're using proto imports
var (
	_ = &fieldmaskpb.FieldMask{}
)

func (c *ServiceChange) IsAdd() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ServiceChange_Added_)
	return ok
}

func (c *ServiceChange) IsModify() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ServiceChange_Modified_)
	return ok
}

func (c *ServiceChange) IsCurrent() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ServiceChange_Current_)
	return ok
}

func (c *ServiceChange) IsDelete() bool {
	if c == nil {
		return false
	}
	_, ok := c.ChangeType.(*ServiceChange_Removed_)
	return ok
}

func (c *ServiceChange) GetCurrentViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ServiceChange_Added_:
		return cType.Added.ViewIndex
	case *ServiceChange_Modified_:
		return cType.Modified.ViewIndex
	}
	return 0
}

func (c *ServiceChange) GetPreviousViewIndex() int32 {
	switch cType := c.ChangeType.(type) {
	case *ServiceChange_Removed_:
		return cType.Removed.ViewIndex
	case *ServiceChange_Modified_:
		return cType.Modified.PreviousViewIndex
	}
	return 0
}

func (c *ServiceChange) GetService() *Service {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ServiceChange_Added_:
		return cType.Added.Service
	case *ServiceChange_Modified_:
		return cType.Modified.Service
	case *ServiceChange_Current_:
		return cType.Current.Service
	case *ServiceChange_Removed_:
		return nil
	}
	return nil
}

func (c *ServiceChange) GetRawResource() gotenresource.Resource {
	return c.GetService()
}

func (c *ServiceChange) GetServiceName() *Name {
	if c == nil {
		return nil
	}
	switch cType := c.ChangeType.(type) {
	case *ServiceChange_Added_:
		return cType.Added.Service.GetName()
	case *ServiceChange_Modified_:
		return cType.Modified.Name
	case *ServiceChange_Current_:
		return cType.Current.Service.GetName()
	case *ServiceChange_Removed_:
		return cType.Removed.Name
	}
	return nil
}

func (c *ServiceChange) GetRawName() gotenresource.Name {
	return c.GetServiceName()
}

func (c *ServiceChange) SetAddedRaw(snapshot gotenresource.Resource, idx int) {
	c.ChangeType = &ServiceChange_Added_{
		Added: &ServiceChange_Added{
			Service:   snapshot.(*Service),
			ViewIndex: int32(idx),
		},
	}
}

func (c *ServiceChange) SetModifiedRaw(name gotenresource.Name, snapshot gotenresource.Resource, prevIdx int, newIdx int) {
	c.ChangeType = &ServiceChange_Modified_{
		Modified: &ServiceChange_Modified{
			Name:              name.(*Name),
			Service:           snapshot.(*Service),
			PreviousViewIndex: int32(prevIdx),
			ViewIndex:         int32(newIdx),
		},
	}
}

func (c *ServiceChange) SetCurrentRaw(snapshot gotenresource.Resource) {
	c.ChangeType = &ServiceChange_Current_{
		Current: &ServiceChange_Current{
			Service: snapshot.(*Service),
		},
	}
}

func (c *ServiceChange) SetDeletedRaw(name gotenresource.Name, idx int) {
	c.ChangeType = &ServiceChange_Removed_{
		Removed: &ServiceChange_Removed{
			Name:      name.(*Name),
			ViewIndex: int32(idx),
		},
	}
}
