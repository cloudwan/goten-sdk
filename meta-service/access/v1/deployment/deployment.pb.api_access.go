// Code generated by protoc-gen-goten-access
// Resource: Deployment
// DO NOT EDIT!!!

package deployment_access

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gotenaccess "github.com/cloudwan/goten-sdk/runtime/access"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
	"github.com/cloudwan/goten-sdk/types/watch_type"

	deployment_client "github.com/cloudwan/goten-sdk/meta-service/client/v1/deployment"
	deployment "github.com/cloudwan/goten-sdk/meta-service/resources/v1/deployment"
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
)

type apiDeploymentAccess struct {
	client deployment_client.DeploymentServiceClient
}

func NewApiDeploymentAccess(client deployment_client.DeploymentServiceClient) deployment.DeploymentAccess {
	return &apiDeploymentAccess{client: client}
}

func (a *apiDeploymentAccess) GetDeployment(ctx context.Context, query *deployment.GetQuery) (*deployment.Deployment, error) {
	if !query.Reference.IsFullyQualified() {
		return nil, status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &deployment_client.GetDeploymentRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	res, err := a.client.GetDeployment(ctx, request)
	if err != nil {
		return nil, err
	}
	query.Reference.Resolve(res)
	return res, nil
}

func (a *apiDeploymentAccess) BatchGetDeployments(ctx context.Context, refs []*deployment.Reference, opts ...gotenresource.BatchGetOption) error {
	batchGetOpts := gotenresource.MakeBatchGetOptions(opts)
	asNames := make([]*deployment.Name, 0, len(refs))
	for _, ref := range refs {
		if !ref.IsFullyQualified() {
			return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
		}
		asNames = append(asNames, &ref.Name)
	}
	request := &deployment_client.BatchGetDeploymentsRequest{
		Names: asNames,
	}
	fieldMask := batchGetOpts.GetFieldMask(deployment.GetDescriptor())
	if fieldMask != nil {
		request.FieldMask = fieldMask.(*deployment.Deployment_FieldMask)
	}
	resp, err := a.client.BatchGetDeployments(ctx, request)
	if err != nil {
		return err
	}
	resultMap := make(map[deployment.Name]*deployment.Deployment, len(refs))
	for _, resolvedRes := range resp.GetDeployments() {
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

func (a *apiDeploymentAccess) QueryDeployments(ctx context.Context, query *deployment.ListQuery) (*deployment.QueryResultSnapshot, error) {
	request := &deployment_client.ListDeploymentsRequest{
		Filter:            query.Filter,
		FieldMask:         query.Mask,
		IncludePagingInfo: query.WithPagingInfo,
	}
	if query.Pager != nil {
		request.PageSize = int32(query.Pager.Limit)
		request.OrderBy = query.Pager.OrderBy
		request.PageToken = query.Pager.Cursor
	}
	resp, err := a.client.ListDeployments(ctx, request)
	if err != nil {
		return nil, err
	}
	return &deployment.QueryResultSnapshot{
		Deployments:       resp.Deployments,
		NextPageCursor:    resp.NextPageToken,
		PrevPageCursor:    resp.PrevPageToken,
		TotalResultsCount: resp.TotalResultsCount,
		CurrentOffset:     resp.CurrentOffset,
	}, nil
}

func (a *apiDeploymentAccess) WatchDeployment(ctx context.Context, query *deployment.GetQuery, observerCb func(*deployment.DeploymentChange) error) error {
	if !query.Reference.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", query.Reference)
	}
	request := &deployment_client.WatchDeploymentRequest{
		Name:      &query.Reference.Name,
		FieldMask: query.Mask,
	}
	changesStream, initErr := a.client.WatchDeployment(ctx, request)
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

func (a *apiDeploymentAccess) WatchDeployments(ctx context.Context, query *deployment.WatchQuery, observerCb func(*deployment.QueryResultChange) error) error {
	request := &deployment_client.WatchDeploymentsRequest{
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
	changesStream, initErr := a.client.WatchDeployments(ctx, request)
	if initErr != nil {
		return initErr
	}
	for {
		respChange, err := changesStream.Recv()
		if err != nil {
			return fmt.Errorf("watch recv error: %w", err)
		}
		changesWithPaging := &deployment.QueryResultChange{
			Changes:      respChange.DeploymentChanges,
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

func (a *apiDeploymentAccess) SaveDeployment(ctx context.Context, res *deployment.Deployment, opts ...gotenresource.SaveOption) error {
	saveOpts := gotenresource.MakeSaveOptions(opts)
	previousRes := saveOpts.GetPreviousResource()

	if previousRes == nil && !saveOpts.OnlyUpdate() && !saveOpts.OnlyCreate() {
		var err error
		previousRes, err = a.GetDeployment(ctx, &deployment.GetQuery{Reference: res.Name.AsReference()})
		if err != nil {
			if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != codes.NotFound {
				return err
			}
		}
	}

	if saveOpts.OnlyUpdate() || previousRes != nil {
		updateRequest := &deployment_client.UpdateDeploymentRequest{
			Deployment: res,
		}
		if updateMask := saveOpts.GetUpdateMask(); updateMask != nil {
			updateRequest.UpdateMask = updateMask.(*deployment.Deployment_FieldMask)
		}
		if mask, conditionalState := saveOpts.GetCAS(); mask != nil && conditionalState != nil {
			updateRequest.Cas = &deployment_client.UpdateDeploymentRequest_CAS{
				ConditionalState: conditionalState.(*deployment.Deployment),
				FieldMask:        mask.(*deployment.Deployment_FieldMask),
			}
		}
		_, err := a.client.UpdateDeployment(ctx, updateRequest)
		if err != nil {
			return err
		}
		return nil
	} else {
		createRequest := &deployment_client.CreateDeploymentRequest{
			Deployment: res,
		}
		_, err := a.client.CreateDeployment(ctx, createRequest)
		if err != nil {
			return err
		}
		return nil
	}
}

func (a *apiDeploymentAccess) DeleteDeployment(ctx context.Context, ref *deployment.Reference, opts ...gotenresource.DeleteOption) error {
	if !ref.IsFullyQualified() {
		return status.Errorf(codes.InvalidArgument, "Reference %s is not fully specified", ref)
	}
	request := &deployment_client.DeleteDeploymentRequest{
		Name: &ref.Name,
	}
	_, err := a.client.DeleteDeployment(ctx, request)
	return err
}

func init() {
	gotenaccess.GetRegistry().RegisterApiAccessConstructor(deployment.GetDescriptor(), func(cc grpc.ClientConnInterface) gotenresource.Access {
		return deployment.AsAnyCastAccess(NewApiDeploymentAccess(deployment_client.NewDeploymentServiceClient(cc)))
	})
}
