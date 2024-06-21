// Code generated by protoc-gen-goten-access
// Resource: Resource
// DO NOT EDIT!!!

package resource_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
	gotenfilter "github.com/cloudwan/goten-sdk/runtime/resource/filter"
	"github.com/cloudwan/goten-sdk/types/watch_type"

	resource_client "github.com/cloudwan/goten-sdk/meta-service/client/v1/resource"
	resource "github.com/cloudwan/goten-sdk/meta-service/resources/v1/resource"
)

var (
	_ = new(context.Context)
	_ = new(fmt.GoStringer)

	_ = new(grpc.ClientConnInterface)
	_ = codes.NotFound
	_ = status.Status{}

	_ = new(gotenaccess.Watcher)
	_ = watch_type.WatchType_STATEFUL
	_ = new(gotenresource.ListQuery)
	_ = gotenfilter.Eq
)

type apiResourceAccess struct {
	client resource_client.ResourceServiceClient
}

func NewApiResourceAccess(client resource_client.ResourceServiceClient) resource.ResourceAccess {
	return &apiResourceAccess{client: client}
}

func (a *apiResourceAccess) GetResource(ctx context.Context, query *resource.GetQuery) (*resource.Resource, error) {
	if !query.Reference.IsFullyQualified() {
		return nil, status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &resource_client.GetResourceRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetResource(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiResourceAccess) BatchGetResources(ctx context.Context, refs []*resource.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	asNames := make([]*resource.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &resource_client.BatchGetResourcesRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(resource.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*resource.Resource_FieldMask)
	}
	resp, err := a.client.BatchGetResources(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[resource.Name]*resource.Resource, len(refs))
	for _, resolvedRes := range resp.GetResources() {
		resultMap[*resolvedRes.GetName()] = resolvedRes
	}
	for _, ref := range refs {
		resolvedRes := resultMap[ref.Name]
		if resolvedRes != nil {
			ref.Resolve(resolvedRes)
		}
	}
	if batchGetOpts.MustResolveAll() && len(resp.GetMissing()) > 0 {
		return status.Errorf(codes.NotFound, "Number of references not found: %d", len(resp.GetMissing()))
	}
	return nil
}

func (a *apiResourceAccess) QueryResources(ctx context.Context, query *resource.ListQuery) (*resource.QueryResultSnapshot, error) {
	request := &resource_client.ListResourcesRequest{
		Filter:            query.Filter,
		FieldMask:         query.Mask,
		IncludePagingInfo: query.WithPagingInfo,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	if query.Filter != nil && query.Filter.GetCondition() != nil {
		request.Filter, request.Parent = getParentAndFilter(query.Filter)
	}
	resp, err := a.client.ListResources(ctx, request)
	if err != nil {
		return nil, err
	}
	return &resource.QueryResultSnapshot{
		Resources:         resp.Resources,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiResourceAccess) WatchResource(ctx context.Context, query *resource.GetQuery, observerCb func(*resource.ResourceChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &resource_client.WatchResourceRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchResource(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		resp, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		change := resp.GetChange()
		if err := observerCb(change); err != nil {
			return err
		}
	}
}

func (a *apiResourceAccess) WatchResources(ctx context.Context, query *resource.WatchQuery, observerCb func(*resource.QueryResultChange) error) error {
	request := &resource_client.WatchResourcesRequest{
		Filter:       query.Filter,
		FieldMask:    query.Mask,
		MaxChunkSize: int32(query.ChunkSize),
		Type:         query.WatchType,
		ResumeToken:  query.ResumeToken,
	}
	if query.Pager != nil {
		request.OrderBy = query.Pager.OrderBy
		request.PageSize = int32(query.Pager.Limit)
		request.PageToken = query.Pager.Cursor
	}
	if query.Filter != nil && query.Filter.GetCondition() != nil {
		request.Filter, request.Parent = getParentAndFilter(query.Filter)
	}
	changesStream, initErr := a.client.WatchResources(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &resource.QueryResultChange{
			Changes:      respChange.ResourceChanges,
			IsCurrent:    respChange.IsCurrent,
			IsHardReset:  respChange.IsHardReset,
			IsSoftReset:  respChange.IsSoftReset,
			ResumeToken:  respChange.ResumeToken,
			SnapshotSize: respChange.SnapshotSize,
		}
		if respChange.PageTokenChange != nil {
			changesWithPaging.PrevPageCursor = respChange.PageTokenChange.PrevPageToken
			changesWithPaging.NextPageCursor = respChange.PageTokenChange.NextPageToken
		}
		if err := observerCb(changesWithPaging); err != nil {
			return err
		}
	}
}

func (a *apiResourceAccess) SaveResource(ctx context.Context, res *resource.Resource, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetResource(ctx, &resource.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}
	var resp *resource.Resource
	var err error
	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &resource_client.UpdateResourceRequest{
			Resource: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*resource.Resource_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &resource_client.UpdateResourceRequest_CAS{
				ConditionalState: conditionalState.(*resource.Resource),
				FieldMask:        mask.(*resource.Resource_FieldMask),
			}
		}
		resp, err = a.client.UpdateResource(ctx, updateRequest)
		if err != nil {
			return err
		}
	} else {
		createRequest := &resource_client.CreateResourceRequest{
			Resource: res,
		}
		resp, err = a.client.CreateResource(ctx, createRequest)
		if err != nil {
			return err
		}
	}
	// Ensure object is updated - but in most shallow way possible
	res.MakeDiffFieldMask(resp).Set(res, resp)
	return nil
}

func (a *apiResourceAccess) DeleteResource(ctx context.Context, ref *resource.Reference, opts ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &resource_client.DeleteResourceRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeleteResource(ctx, request)
	return err
}
func getParentAndFilter(fullFilter *resource.Filter) (*resource.Filter, *resource.ParentName) {
	var withParentExtraction func(cnd resource.FilterCondition) resource.FilterCondition
	var resultParent *resource.ParentName
	var resultFilter *resource.Filter
	withParentExtraction = func(cnd resource.FilterCondition) resource.FilterCondition {
		switch tCnd := cnd.(type) {
		case *resource.FilterConditionComposite:
			if tCnd.GetOperator() == gotenfilter.AND {
				withoutParentCnds := make([]resource.FilterCondition, 0)
				for _, subCnd := range tCnd.Conditions {
					if subCndNoParent := withParentExtraction(subCnd); subCndNoParent != nil {
						withoutParentCnds = append(withoutParentCnds, subCndNoParent)
					}
				}
				if len(withoutParentCnds) == 0 {
					return nil
				}
				if len(withoutParentCnds) == 1 {
					return withoutParentCnds[0]
				}
				return resource.AndFilterConditions(withoutParentCnds...)
			} else {
				return tCnd
			}
		case *resource.FilterConditionCompare:
			if tCnd.GetOperator() == gotenfilter.Eq && tCnd.GetRawFieldPath().String() == "name" {
				nameValue := tCnd.GetRawValue().(*resource.Name)
				if nameValue != nil && nameValue.ParentName.IsSpecified() {
					resultParent = &nameValue.ParentName
					if nameValue.IsFullyQualified() {
						return tCnd
					}
					return nil
				}
			}
			return tCnd
		default:
			return tCnd
		}
	}
	cndWithoutParent := withParentExtraction(fullFilter.GetCondition())
	if cndWithoutParent != nil {
		resultFilter = &resource.Filter{FilterCondition: cndWithoutParent}
	}
	return resultFilter, resultParent
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(resource.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return resource.AsAnyCastAccess(NewApiResourceAccess(resource_client.NewResourceServiceClient(cc)))
	})
}
