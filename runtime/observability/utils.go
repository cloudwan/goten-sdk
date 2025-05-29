package observability

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func StoreHeadersForRoutedCall(ctx context.Context, md metadata.MD) {
	t := CallTrackerFromContext(ctx)
	if t != nil {
		md.Set(StdUserAgentHeader, t.GetUserAgent())
		md.Set(GotenRequestIdHeader, t.GetCallId())
		md.Set(GotenOriginalIpAddress, t.GetCallerIPAddress())
		md.Set(GotenRoutingRegionHeader, t.GetRoutingRegionId())
		md.Set(GotenExecutingRegionsHeader, t.GetExecutingRegionIds()...)
	}
}
