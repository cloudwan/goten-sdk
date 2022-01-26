package clipb

import (
	"context"

	"google.golang.org/grpc"
)

type GrpcDialer func(ctx context.Context, dialOpts *DialOptions, extra ...grpc.DialOption) (*grpc.ClientConn, error)
