package observability

import (
	"context"
	"slices"
	"sync"
	"sync/atomic"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/cloudwan/goten-sdk/runtime/client"
)

type GrpcServerObserver interface {
	OnNewRequestStarted(ctx context.Context, tracker *CallTracker, st *stats.Begin)
	OnRequestCompleted(ctx context.Context, tracker *CallTracker, st *stats.End)
	OnClientMsgReceived(ctx context.Context, request proto.Message, tracker *CallTracker, st *stats.InPayload)
	OnServerMessageSent(ctx context.Context, response proto.Message, tracker *CallTracker, st *stats.OutPayload)
}

type IdleGrpcServerObserver struct{}

func (*IdleGrpcServerObserver) OnNewRequestStarted(context.Context, *CallTracker, *stats.Begin) {}
func (*IdleGrpcServerObserver) OnRequestCompleted(context.Context, *CallTracker, *stats.End)    {}
func (*IdleGrpcServerObserver) OnClientMsgReceived(context.Context, proto.Message, *CallTracker, *stats.InPayload) {
}
func (*IdleGrpcServerObserver) OnServerMessageSent(context.Context, proto.Message, *CallTracker, *stats.OutPayload) {
}

type GrpcServerStatsHandlerCfg struct {
	TrackKSlowestUnaryRequests  int
	TrackKSlowestStreamRequests int
	TrackKFailedRequests        int
	IpExtractFunction           func(ctx context.Context, md metadata.MD) string
}

type GrpcServerStats struct {
	ConnsActiveCounter  int64
	UnaryActiveCounter  int64
	StreamActiveCounter int64
	ConnsCumCounter     int64
	UnaryCumCounter     int64
	StreamCumCounter    int64
}

type GrpcServerStatsHandler struct {
	cfg           GrpcServerStatsHandlerCfg
	lock          sync.RWMutex
	callObservers []GrpcServerObserver
	slowestCalls  map[client.MethodDescriptor][]*CallTracker
	failedCalls   map[client.MethodDescriptor]map[codes.Code][]*CallTracker

	connsActiveCounter  int64
	unaryActiveCounter  int64
	streamActiveCounter int64
	connsCumCounter     int64
	unaryCumCounter     int64
	streamCumCounter    int64
}

func NewGrpcStatsHandler(cfg GrpcServerStatsHandlerCfg) *GrpcServerStatsHandler {
	return &GrpcServerStatsHandler{
		cfg:          cfg,
		slowestCalls: make(map[client.MethodDescriptor][]*CallTracker),
		failedCalls:  make(map[client.MethodDescriptor]map[codes.Code][]*CallTracker),
	}
}

func (h *GrpcServerStatsHandler) RegisterObserver(o GrpcServerObserver) {
	h.lock.Lock()
	h.callObservers = append(h.callObservers, o)
	h.lock.Unlock()
}

func (h *GrpcServerStatsHandler) GetCurrentStats() GrpcServerStats {
	return GrpcServerStats{
		ConnsActiveCounter:  atomic.LoadInt64(&h.connsActiveCounter),
		UnaryActiveCounter:  atomic.LoadInt64(&h.unaryActiveCounter),
		StreamActiveCounter: atomic.LoadInt64(&h.streamActiveCounter),
		ConnsCumCounter:     atomic.LoadInt64(&h.connsCumCounter),
		UnaryCumCounter:     atomic.LoadInt64(&h.unaryCumCounter),
		StreamCumCounter:    atomic.LoadInt64(&h.streamCumCounter),
	}
}

func (h *GrpcServerStatsHandler) GetCurrentSlowestCalls(md client.MethodDescriptor) []*CallTracker {
	h.lock.RLock()
	slowest := append(make([]*CallTracker, 0), h.slowestCalls[md]...)
	h.lock.RUnlock()
	return slowest
}

func (h *GrpcServerStatsHandler) GetLastFailedCalls(md client.MethodDescriptor, code codes.Code) []*CallTracker {
	h.lock.RLock()
	failed := append(make([]*CallTracker, 0), h.failedCalls[md][code]...)
	h.lock.RUnlock()
	return failed
}

func (h *GrpcServerStatsHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	md := client.GetRegistry().FindMethodDescriptor(info.FullMethodName)
	if md != nil {
		var tracker *CallTracker
		var rootSpan *Span
		ctx, tracker = WithCallTracker(ctx, md)
		ctx, rootSpan = StartRootSpan(ctx, info.FullMethodName)
		tracker.SetRootSpan(rootSpan)
	}
	return ctx
}

func (h *GrpcServerStatsHandler) HandleRPC(ctx context.Context, st stats.RPCStats) {
	switch tSt := st.(type) {
	case *stats.InHeader:
		tracker := CallTrackerFromContext(ctx)
		if tracker != nil {
			tracker.HandleInHeader(tSt)
			if h.cfg.IpExtractFunction != nil {
				tracker.originalIp = h.cfg.IpExtractFunction(ctx, tracker.clientHeader)
			} else {
				tracker.originalIp = extractCallerIp(tracker.clientHeader)
			}
			if (tracker.originalIp == "" || tracker.originalIp == UnknownIPAddress) && tSt.RemoteAddr != nil {
				tracker.originalIp = tSt.RemoteAddr.String()
			}
		}
	case *stats.Begin:
		tracker := CallTrackerFromContext(ctx)
		if tracker != nil {
			tracker.HandleRPCStatsBegin(tSt)
			if tracker.methodDesc.IsUnary() {
				atomic.AddInt64(&h.unaryActiveCounter, 1)
				atomic.AddInt64(&h.unaryCumCounter, 1)
			} else {
				atomic.AddInt64(&h.streamActiveCounter, 1)
				atomic.AddInt64(&h.streamCumCounter, 1)
			}
			for _, observer := range h.callObservers {
				observer.OnNewRequestStarted(ctx, tracker, tSt)
			}
		}
	case *stats.InPayload:
		tracker := CallTrackerFromContext(ctx)
		if tracker != nil {
			msg, ok := tSt.Payload.(proto.Message)
			if ok {
				tracker.HandleRPCInPayload(tSt)
				for _, observer := range h.callObservers {
					observer.OnClientMsgReceived(ctx, msg, tracker, tSt)
				}
			}
		}
	case *stats.OutPayload:
		tracker := CallTrackerFromContext(ctx)
		if tracker != nil {
			msg, ok := tSt.Payload.(proto.Message)
			if ok {
				tracker.HandleRPCOutPayload(tSt)
				for _, observer := range h.callObservers {
					observer.OnServerMessageSent(ctx, msg, tracker, tSt)
				}
			}
			// for streaming requests, span ends with first response sent.
			// TODO: Add configurability, like use CallTracker, get root span, disable
			// auto-end function?
			if !tracker.methodDesc.IsUnary() {
				tracker.rootSpan.End(nil)
				h.recordSlowestReq(tracker, h.cfg.TrackKSlowestStreamRequests)
			}
		}
	// TODO: Handle these 2? Is it needed?
	case *stats.OutHeader:
	case *stats.OutTrailer:
	case *stats.End:
		tracker := CallTrackerFromContext(ctx)
		if tracker != nil {
			tracker.HandleRPCStatsEnd(tSt)
			if tracker.methodDesc.IsUnary() {
				atomic.AddInt64(&h.unaryActiveCounter, -1)
			} else {
				atomic.AddInt64(&h.streamActiveCounter, -1)
			}
			for _, observer := range h.callObservers {
				observer.OnRequestCompleted(ctx, tracker, tSt)
			}
			if tSt.Error != nil && h.cfg.TrackKFailedRequests > 0 {
				h.recordFailedCall(tracker)
			} else if tSt.Error == nil && tracker.methodDesc.IsUnary() {
				h.recordSlowestReq(tracker, h.cfg.TrackKSlowestUnaryRequests)
			}
		}
	}
}

func (h *GrpcServerStatsHandler) TagConn(ctx context.Context, _ *stats.ConnTagInfo) context.Context {
	return ctx
}

func (h *GrpcServerStatsHandler) HandleConn(_ context.Context, st stats.ConnStats) {
	switch st.(type) {
	case *stats.ConnBegin:
		atomic.AddInt64(&h.connsActiveCounter, 1)
		atomic.AddInt64(&h.connsCumCounter, 1)
	case *stats.ConnEnd:
		atomic.AddInt64(&h.connsActiveCounter, -1)
	}
}

func (h *GrpcServerStatsHandler) recordFailedCall(tracker *CallTracker) {
	code := status.Code(tracker.exitErr)

	h.lock.Lock()
	byMethod := h.failedCalls[tracker.methodDesc]
	if byMethod == nil {
		byMethod = make(map[codes.Code][]*CallTracker)
		h.failedCalls[tracker.methodDesc] = byMethod
	}
	byCode := byMethod[code]
	if len(byCode) == 0 {
		byCode = make([]*CallTracker, 0, h.cfg.TrackKFailedRequests)
	} else if len(byCode) == h.cfg.TrackKFailedRequests {
		byCode = slices.Delete(byCode, 0, 1)
	}
	byCode = append(byCode, tracker)
	byMethod[code] = byCode
	h.lock.Unlock()
}

func (h *GrpcServerStatsHandler) recordSlowestReq(tracker *CallTracker, maxToKeep int) {
	if maxToKeep <= 0 {
		return
	}
	h.lock.RLock()
	currentSlowest := h.slowestCalls[tracker.methodDesc]
	addThisCall := len(currentSlowest) < maxToKeep
	if !addThisCall {
		thisDuration := tracker.rootSpan.end.Sub(tracker.rootSpan.begin)
		otherDuration := currentSlowest[0].rootSpan.end.Sub(currentSlowest[0].rootSpan.begin)
		addThisCall = thisDuration > otherDuration
	}
	h.lock.RUnlock()

	if addThisCall {
		h.lock.Lock()
		updatedSlowest := h.slowestCalls[tracker.methodDesc]
		if len(updatedSlowest) > maxToKeep {
			updatedSlowest = slices.Delete(updatedSlowest, 0, 1)
		}
		updatedSlowest = append(updatedSlowest, tracker)
		slices.SortFunc(updatedSlowest, func(a, b *CallTracker) int {
			thisDuration := a.endTime.Sub(a.beginTime)
			otherDuration := b.endTime.Sub(b.beginTime)
			if thisDuration < otherDuration {
				return -1
			} else if thisDuration == otherDuration {
				return 0
			}
			return 1
		})
		h.slowestCalls[tracker.methodDesc] = updatedSlowest
		h.lock.Unlock()
	}
}
