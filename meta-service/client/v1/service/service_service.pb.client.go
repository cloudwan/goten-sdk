// Code generated by protoc-gen-goten-client
// API: ServiceService
// DO NOT EDIT!!!

package service_client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	service "github.com/cloudwan/goten-sdk/meta-service/resources/v1/service"
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
	_ = &service.Service{}
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

// ServiceServiceClient is the client API for ServiceService.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceServiceClient interface {
	GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*service.Service, error)
	BatchGetServices(ctx context.Context, in *BatchGetServicesRequest, opts ...grpc.CallOption) (*BatchGetServicesResponse, error)
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
	WatchService(ctx context.Context, in *WatchServiceRequest, opts ...grpc.CallOption) (WatchServiceClientStream, error)
	WatchServices(ctx context.Context, in *WatchServicesRequest, opts ...grpc.CallOption) (WatchServicesClientStream, error)
	CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*service.Service, error)
	UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*service.Service, error)
	DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type client struct {
	cc grpc.ClientConnInterface
}

func NewServiceServiceClient(cc grpc.ClientConnInterface) ServiceServiceClient {
	return &client{cc}
}

func (c *client) GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*service.Service, error) {
	out := new(service.Service)
	err := c.cc.Invoke(ctx, "/goten.meta.v1.ServiceService/GetService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) BatchGetServices(ctx context.Context, in *BatchGetServicesRequest, opts ...grpc.CallOption) (*BatchGetServicesResponse, error) {
	out := new(BatchGetServicesResponse)
	err := c.cc.Invoke(ctx, "/goten.meta.v1.ServiceService/BatchGetServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, "/goten.meta.v1.ServiceService/ListServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) WatchService(ctx context.Context, in *WatchServiceRequest, opts ...grpc.CallOption) (WatchServiceClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchService",
			ServerStreams: true,
		},
		"/goten.meta.v1.ServiceService/WatchService", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchServiceWatchServiceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchServiceClientStream interface {
	Recv() (*WatchServiceResponse, error)
	grpc.ClientStream
}

type watchServiceWatchServiceClient struct {
	grpc.ClientStream
}

func (x *watchServiceWatchServiceClient) Recv() (*WatchServiceResponse, error) {
	m := new(WatchServiceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) WatchServices(ctx context.Context, in *WatchServicesRequest, opts ...grpc.CallOption) (WatchServicesClientStream, error) {
	stream, err := c.cc.NewStream(ctx,
		&grpc.StreamDesc{
			StreamName:    "WatchServices",
			ServerStreams: true,
		},
		"/goten.meta.v1.ServiceService/WatchServices", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchServicesWatchServicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchServicesClientStream interface {
	Recv() (*WatchServicesResponse, error)
	grpc.ClientStream
}

type watchServicesWatchServicesClient struct {
	grpc.ClientStream
}

func (x *watchServicesWatchServicesClient) Recv() (*WatchServicesResponse, error) {
	m := new(WatchServicesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*service.Service, error) {
	out := new(service.Service)
	err := c.cc.Invoke(ctx, "/goten.meta.v1.ServiceService/CreateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*service.Service, error) {
	out := new(service.Service)
	err := c.cc.Invoke(ctx, "/goten.meta.v1.ServiceService/UpdateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/goten.meta.v1.ServiceService/DeleteService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}