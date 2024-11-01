// Code generated by protoc-gen-goten-client
// API: DeploymentService
// DO NOT EDIT!!!

package deployment_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	deployment "github.com/cloudwan/goten-sdk/meta-service/resources/v1/deployment"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = protoimpl.DescBuilder{}
	_ = new(context.Context)
	_ = grpc.ClientConn{}
)

// make sure we're using proto imports
var (
	_ = &emptypb.Empty{}
	_ = &deployment.Deployment{}
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeploymentServiceClient is the client API for DeploymentService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeploymentServiceClient interface {
	GetDeployment(ctx context.Context, in *GetDeploymentRequest, opts ...grpc.CallOption) (*deployment.Deployment, error)
	BatchGetDeployments(ctx context.Context, in *BatchGetDeploymentsRequest, opts ...grpc.CallOption) (*BatchGetDeploymentsResponse, error)
	ListDeployments(ctx context.Context, in *ListDeploymentsRequest, opts ...grpc.CallOption) (*ListDeploymentsResponse, error)
	WatchDeployment(ctx context.Context, in *WatchDeploymentRequest, opts ...grpc.CallOption) (WatchDeploymentClientStream, error)
	WatchDeployments(ctx context.Context, in *WatchDeploymentsRequest, opts ...grpc.CallOption) (WatchDeploymentsClientStream, error)
	CreateDeployment(ctx context.Context, in *CreateDeploymentRequest, opts ...grpc.CallOption) (*deployment.Deployment, error)
	UpdateDeployment(ctx context.Context, in *UpdateDeploymentRequest, opts ...grpc.CallOption) (*deployment.Deployment, error)
	DeleteDeployment(ctx context.Context, in *DeleteDeploymentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BeginUpgrade(ctx context.Context, in *BeginUpgradeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetAvailableUpgrade(ctx context.Context, in *SetAvailableUpgradeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	NotifyShardsUpgradeReadiness(ctx context.Context, in *NotifyShardsUpgradeReadinessRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetOrRegisterDataUpdateTask(ctx context.Context, in *GetOrRegisterDataUpdateTaskRequest, opts ...grpc.CallOption) (*GetOrRegisterDataUpdateTaskResponse, error)
	UpdateDataTaskUpdate(ctx context.Context, in *UpdateDataTaskUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewDeploymentServiceClient(cc grpc.ClientConnInterface) DeploymentServiceClient {
	return &client{cc}
}

func (c *client) GetDeployment(ctx context.Context, in *GetDeploymentRequest, opts ...grpc.CallOption) (*deployment.Deployment, error) {
	out := new(deployment.Deployment)
	err := c.cc.Invoke(ctx, "/goten.meta.v1.DeploymentService/GetDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetDeployments(ctx context.Context, in *BatchGetDeploymentsRequest, opts ...grpc.CallOption) (*BatchGetDeploymentsResponse, error) {
	out := new(BatchGetDeploymentsResponse)
	err := c.cc.Invoke(ctx, "/goten.meta.v1.DeploymentService/BatchGetDeployments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListDeployments(ctx context.Context, in *ListDeploymentsRequest, opts ...grpc.CallOption) (*ListDeploymentsResponse, error) {
	out := new(ListDeploymentsResponse)
	err := c.cc.Invoke(ctx, "/goten.meta.v1.DeploymentService/ListDeployments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchDeployment(ctx context.Context, in *WatchDeploymentRequest, opts ...grpc.CallOption) (WatchDeploymentClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchDeployment",
			ServerStreams: true,
		},
		"/goten.meta.v1.DeploymentService/WatchDeployment", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchDeploymentWatchDeploymentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchDeploymentClientStream interface {
	Recv() (*WatchDeploymentResponse, error)
	grpc.ClientStream
}

type watchDeploymentWatchDeploymentClient struct {
	grpc.ClientStream
}

func (x *watchDeploymentWatchDeploymentClient) Recv() (*WatchDeploymentResponse, error) {
	m := new(WatchDeploymentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchDeployments(ctx context.Context, in *WatchDeploymentsRequest, opts ...grpc.CallOption) (WatchDeploymentsClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchDeployments",
			ServerStreams: true,
		},
		"/goten.meta.v1.DeploymentService/WatchDeployments", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchDeploymentsWatchDeploymentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchDeploymentsClientStream interface {
	Recv() (*WatchDeploymentsResponse, error)
	grpc.ClientStream
}

type watchDeploymentsWatchDeploymentsClient struct {
	grpc.ClientStream
}

func (x *watchDeploymentsWatchDeploymentsClient) Recv() (*WatchDeploymentsResponse, error) {
	m := new(WatchDeploymentsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateDeployment(ctx context.Context, in *CreateDeploymentRequest, opts ...grpc.CallOption) (*deployment.Deployment, error) {
	out := new(deployment.Deployment)
	err := c.cc.Invoke(ctx, "/goten.meta.v1.DeploymentService/CreateDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateDeployment(ctx context.Context, in *UpdateDeploymentRequest, opts ...grpc.CallOption) (*deployment.Deployment, error) {
	out := new(deployment.Deployment)
	err := c.cc.Invoke(ctx, "/goten.meta.v1.DeploymentService/UpdateDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteDeployment(ctx context.Context, in *DeleteDeploymentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/goten.meta.v1.DeploymentService/DeleteDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BeginUpgrade(ctx context.Context, in *BeginUpgradeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/goten.meta.v1.DeploymentService/BeginUpgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) SetAvailableUpgrade(ctx context.Context, in *SetAvailableUpgradeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/goten.meta.v1.DeploymentService/SetAvailableUpgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) NotifyShardsUpgradeReadiness(ctx context.Context, in *NotifyShardsUpgradeReadinessRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/goten.meta.v1.DeploymentService/NotifyShardsUpgradeReadiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) GetOrRegisterDataUpdateTask(ctx context.Context, in *GetOrRegisterDataUpdateTaskRequest, opts ...grpc.CallOption) (*GetOrRegisterDataUpdateTaskResponse, error) {
	out := new(GetOrRegisterDataUpdateTaskResponse)
	err := c.cc.Invoke(ctx, "/goten.meta.v1.DeploymentService/GetOrRegisterDataUpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateDataTaskUpdate(ctx context.Context, in *UpdateDataTaskUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/goten.meta.v1.DeploymentService/UpdateDataTaskUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
