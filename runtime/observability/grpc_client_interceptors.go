package observability

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// RequestIdUnaryClientInterceptor carries the request ID without forwarding
// incoming credentials or changing existing outgoing metadata.
func RequestIdUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	return invoker(withRequestIdOutgoingContext(ctx), method, req, reply, cc, opts...)
}

// RequestIdStreamClientInterceptor carries the request ID when opening a stream.
func RequestIdStreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return streamer(withRequestIdOutgoingContext(ctx), desc, cc, method, opts...)
}

func withRequestIdOutgoingContext(ctx context.Context) context.Context {
	tracker := CallTrackerFromContext(ctx)
	if tracker == nil {
		return ctx
	}
	md, _ := metadata.FromOutgoingContext(ctx)
	if md == nil {
		md = metadata.MD{}
	}
	md.Set(GotenRequestIdHeader, tracker.GetRequestId())
	return metadata.NewOutgoingContext(ctx, md)
}
