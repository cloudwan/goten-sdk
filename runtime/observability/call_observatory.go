package observability

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"github.com/cloudwan/goten-sdk/runtime/client"
)

var (
	callObservers []CallObserver
)

// DEPRECATED: Use GrpcServerStatsHandler and register observers there.
type CallObserver interface {
	OnUnaryCallBegan(ctx context.Context, request proto.Message, tracker *CallTracker)
	OnUnaryCallCompleted(ctx context.Context, response proto.Message, tracker *CallTracker)
	OnStreamingCallBegan(stream grpc.ServerStream, tracker *CallTracker)
	OnStreamingCallCompleted(stream grpc.ServerStream, tracker *CallTracker)
	OnStreamingMsgReceived(stream grpc.ServerStream, request proto.Message, tracker *CallTracker)
	OnStreamingMsgSent(stream grpc.ServerStream, response proto.Message, tracker *CallTracker)
}

// Not exactly thread safe, but its not like we need it now?

// DEPRECATED: Use GrpcServerStatsHandler and register observers there.
func RegisterCallObserver(observer CallObserver) {
	callObservers = append(callObservers, observer)
}

func UnRegisterCallObserver(observer CallObserver) {
	newObservers := make([]CallObserver, 0)
	for _, current := range callObservers {
		if current != observer {
			newObservers = append(newObservers, current)
		}
	}
	callObservers = newObservers
}

// DEPRECATED: Use GrpcServerStatsHandler
func WithUnaryObserving(ctx context.Context, request proto.Message, methodDesc client.MethodDescriptor) context.Context {
	// detect if new-style tracker is here... if not, then we have legacy app
	trackerInterface := ctx.Value(callTrackerContextKey)
	var tracker *CallTracker
	if trackerInterface == nil {
		tracker = newCallTrackerLegacy(ctx, methodDesc)
		ctx = context.WithValue(ctx, callTrackerContextKey, tracker)
	} else {
		tracker = trackerInterface.(*CallTracker)
	}
	// legacy observers
	for _, observer := range callObservers {
		observer.OnUnaryCallBegan(ctx, request, tracker)
	}
	return ctx
}

// DEPRECATED: Use GrpcServerStatsHandler
func NotifyUnaryCallCompleted(ctx context.Context, response proto.Message, err error) {
	trackerInterface := ctx.Value(callTrackerContextKey)
	if trackerInterface == nil {
		return
	}
	tracker := trackerInterface.(*CallTracker)
	if tracker.isLegacy {
		tracker.endTime = time.Now().UTC()
		tracker.exitErr = err
	}
	// legacy observers
	for _, observer := range callObservers {
		observer.OnUnaryCallCompleted(ctx, response, tracker)
	}
}

type observableStream struct {
	grpc.ServerStream
	ctx     context.Context
	tracker *CallTracker
}

// DEPRECATED: Use GrpcServerStatsHandler
func WithObservingStream(inner grpc.ServerStream, methodDesc client.MethodDescriptor) grpc.ServerStream {
	ctx := inner.Context()
	trackerInterface := ctx.Value(callTrackerContextKey)
	var tracker *CallTracker
	if trackerInterface == nil {
		tracker = newCallTrackerLegacy(ctx, methodDesc)
		ctx = context.WithValue(ctx, callTrackerContextKey, tracker)
	} else {
		tracker = trackerInterface.(*CallTracker)
	}
	stream := &observableStream{
		ServerStream: inner,
		ctx:          ctx,
		tracker:      tracker,
	}
	// legacy observers
	for _, observer := range callObservers {
		observer.OnStreamingCallBegan(stream, tracker)
	}
	return stream
}

// DEPRECATED: Use GrpcServerStatsHandler
func NotifyStreamCallCompleted(stream grpc.ServerStream, err error) {
	trackerInterface := stream.Context().Value(callTrackerContextKey)
	if trackerInterface == nil {
		return
	}
	tracker := trackerInterface.(*CallTracker)
	if tracker.isLegacy {
		tracker.endTime = time.Now().UTC()
		tracker.exitErr = err
	}
	// legacy observers
	for _, observer := range callObservers {
		observer.OnStreamingCallCompleted(stream, tracker)
	}
}

func (os *observableStream) Context() context.Context {
	return os.ctx
}

func (os *observableStream) SendMsg(m interface{}) error {
	err := os.ServerStream.SendMsg(m)
	if err == nil {
		protoMsg := m.(proto.Message)
		for _, observer := range callObservers {
			observer.OnStreamingMsgSent(os, protoMsg, os.tracker)
		}
	}
	return err
}

func (os *observableStream) RecvMsg(m interface{}) error {
	err := os.ServerStream.RecvMsg(m)
	if err == nil {
		protoMsg := m.(proto.Message)
		for _, observer := range callObservers {
			observer.OnStreamingMsgReceived(os, protoMsg, os.tracker)
		}
	}
	return err
}
