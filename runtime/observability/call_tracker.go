package observability

import (
	"context"
	"slices"
	"strings"
	"time"

	"github.com/lithammer/shortuuid/v4"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/stats"
	"google.golang.org/protobuf/proto"

	"github.com/cloudwan/goten-sdk/runtime/client"
	"github.com/cloudwan/goten-sdk/runtime/resource"
)

const (
	UnknownIPAddress = "UnknownIP"
	UnknownUserAgent = "Unknown"
)

const (
	// standard HTTP headers used by tracker

	StdUserAgentHeader    = "user-agent"
	StdRealIpHeader       = "x-real-ip"
	StdForwardedForHeader = "x-forwarded-for"

	// Goten specific HTTP headers used by tracker

	GotenRequestIdHeader        = "x-goten-request-id"
	GotenOriginalIpAddress      = "x-goten-original-ip"
	GotenRoutingRegionHeader    = "x-goten-routing-region"
	GotenExecutingRegionsHeader = "x-goten-executing-regions"

	MaxRememberedTxRecordsInCallTracker = 10
)

type callTrackerContextKeyType struct{}

var callTrackerContextKey callTrackerContextKeyType

func CallTrackerFromContext(ctx context.Context) *CallTracker {
	tracker := ctx.Value(callTrackerContextKey)
	if tracker != nil {
		return tracker.(*CallTracker)
	}
	return nil
}

type CallTracker struct {
	clientHeader      metadata.MD
	callId            string
	origCallId        string
	methodDesc        client.MethodDescriptor
	beginTime         time.Time
	endTime           time.Time
	exitErr           error
	originalIp        string
	userAgent         string
	executedViaRegion string
	executedByRegions []string
	customLabels      map[string]string
	customAttrs       map[string]interface{}
	txTracker         *TxTracker
	failedTxs         []*TxTracker
	successfulTxs     []*TxTracker
	rootSpan          *Span
	isLegacy          bool
	origInPayload     *stats.InPayload
	firstOutPayload   *stats.OutPayload
}

func WithCallTracker(ctx context.Context, methodDesc client.MethodDescriptor) (context.Context, *CallTracker) {
	tracker := &CallTracker{
		callId:       shortuuid.New(),
		methodDesc:   methodDesc,
		customLabels: map[string]string{},
		customAttrs:  map[string]interface{}{},
		originalIp:   UnknownIPAddress,
		userAgent:    UnknownUserAgent,
	}
	ctx = context.WithValue(ctx, callTrackerContextKey, tracker)
	return ctx, tracker
}

func newCallTrackerLegacy(ctx context.Context, methodDesc client.MethodDescriptor) *CallTracker {
	tracker := &CallTracker{
		callId:       shortuuid.New(),
		beginTime:    time.Now().UTC(),
		methodDesc:   methodDesc,
		customLabels: map[string]string{},
		customAttrs:  map[string]interface{}{},
		originalIp:   UnknownIPAddress,
		userAgent:    UnknownUserAgent,
		isLegacy:     true,
	}

	// populate some data from HTTP headers
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		tracker.clientHeader = md
		userAgentVals := md.Get(StdUserAgentHeader)
		if len(userAgentVals) > 0 {
			tracker.userAgent = userAgentVals[0]
		}

		tracker.originalIp = extractCallerIp(md)
		if tracker.originalIp == UnknownIPAddress {
			p, ok := peer.FromContext(ctx)
			if ok && p != nil {
				tracker.originalIp = p.Addr.String()
			}
		}

		viaRegionVals := md.Get(GotenRoutingRegionHeader)
		if len(viaRegionVals) > 0 {
			tracker.executedViaRegion = viaRegionVals[0]
		}

		fromRegionVals := md.Get(GotenExecutingRegionsHeader)
		tracker.executedByRegions = fromRegionVals

		callIdVals := md.Get(GotenRequestIdHeader)
		if len(callIdVals) > 0 {
			tracker.origCallId = callIdVals[0]
		}
	}
	return tracker
}

func (t *CallTracker) HandleInHeader(st *stats.InHeader) {
	t.clientHeader = make(metadata.MD, len(st.Header))
	for key, values := range st.Header {
		t.clientHeader[strings.ToLower(key)] = values
	}
	userAgentVals := t.clientHeader.Get(StdUserAgentHeader)
	if len(userAgentVals) > 0 {
		t.userAgent = userAgentVals[0]
	}

	viaRegionVals := t.clientHeader.Get(GotenRoutingRegionHeader)
	if len(viaRegionVals) > 0 {
		t.executedViaRegion = viaRegionVals[0]
	}

	fromRegionVals := t.clientHeader.Get(GotenExecutingRegionsHeader)
	t.executedByRegions = fromRegionVals

	callIdVals := t.clientHeader.Get(GotenRequestIdHeader)
	if len(callIdVals) > 0 {
		t.origCallId = callIdVals[0]
	}
}

func (t *CallTracker) HandleRPCStatsBegin(st *stats.Begin) {
	t.beginTime = st.BeginTime
}

func (t *CallTracker) HandleRPCStatsEnd(st *stats.End) {
	t.endTime = st.EndTime
	t.exitErr = st.Error
	// streaming spans are closed earlier...
	if t.methodDesc.IsUnary() {
		t.rootSpan.EndWithTime(st.Error, st.EndTime)
	}
}

func (t *CallTracker) HandleRPCInPayload(st *stats.InPayload) {
	if t != nil && t.origInPayload == nil {
		t.origInPayload = st
	}
}

func (t *CallTracker) HandleRPCOutPayload(st *stats.OutPayload) {
	if t != nil && t.firstOutPayload == nil {
		t.firstOutPayload = st
	}
}

func (t *CallTracker) SetRootSpan(span *Span) {
	if t != nil {
		t.rootSpan = span
	}
}

func (t *CallTracker) SetCallerIpAddress(ip string) {
	if t != nil {
		t.originalIp = ip
	}
}

func (t *CallTracker) GetOriginalPayloadStats() *stats.InPayload {
	if t == nil {
		return nil
	}
	return t.origInPayload
}

func (t *CallTracker) GetOriginalRequest() proto.Message {
	if t == nil || t.origInPayload == nil {
		return nil
	}
	return t.origInPayload.Payload.(proto.Message)
}

func (t *CallTracker) GetFirstOutPayloadStats() *stats.OutPayload {
	if t == nil {
		return nil
	}
	return t.firstOutPayload
}

func (t *CallTracker) GetFirstResponse() proto.Message {
	if t == nil || t.firstOutPayload == nil {
		return nil
	}
	return t.firstOutPayload.Payload.(proto.Message)
}

func (t *CallTracker) GetClientHeaders() metadata.MD {
	if t == nil {
		return nil
	}
	return t.clientHeader
}

func (t *CallTracker) GetCurrentMethod() client.MethodDescriptor {
	if t == nil {
		return nil
	}
	return t.methodDesc
}

func (t *CallTracker) GetCurrentTx() *TxTracker {
	if t == nil {
		return nil
	}
	return t.txTracker
}

func (t *CallTracker) GetSuccessfulTxs() []*TxTracker {
	if t == nil {
		return nil
	}
	return t.successfulTxs
}

func (t *CallTracker) GetFailedTxs() []*TxTracker {
	if t == nil {
		return nil
	}
	return t.failedTxs
}

func (t *CallTracker) GetCallId() string {
	if t == nil {
		return ""
	}
	return t.callId
}

func (t *CallTracker) GetExitErr() error {
	if t == nil {
		return nil
	}
	return t.exitErr
}

func (t *CallTracker) GetBeginTime() time.Time {
	if t == nil {
		return time.Time{}
	}
	return t.beginTime
}

func (t *CallTracker) GetEndTime() time.Time {
	if t == nil {
		return time.Time{}
	}
	return t.endTime
}

func (t *CallTracker) GetOrigId() string {
	if t == nil {
		return ""
	}
	return t.origCallId
}

func (t *CallTracker) GetUserAgent() string {
	if t == nil {
		return ""
	}
	return t.userAgent
}

func (t *CallTracker) GetCallerIPAddress() string {
	if t == nil {
		return ""
	}
	return t.originalIp
}

func (t *CallTracker) GetRoutingRegionId() string {
	if t == nil {
		return ""
	}
	return t.executedViaRegion
}

func (t *CallTracker) GetExecutingRegionIds() []string {
	if t == nil {
		return nil
	}
	return t.executedByRegions
}

func (t *CallTracker) GetLabel(key string) string {
	if t != nil {
		return t.customLabels[key]
	}
	return ""
}

func (t *CallTracker) GetRootSpan() *Span {
	if t == nil {
		return nil
	}
	return t.rootSpan
}

func (t *CallTracker) GetStringAttribute(key string) string {
	if t != nil {
		attr := t.customAttrs[key]
		if attr == nil {
			return ""
		}
		return attr.(string)
	}
	return ""
}

func (t *CallTracker) GetStringsArrayAttribute(key string) []string {
	if t != nil {
		attr := t.customAttrs[key]
		if attr == nil {
			return nil
		}
		return attr.([]string)
	}
	return nil
}

func (t *CallTracker) GetAttribute(key string) interface{} {
	if t != nil {
		return t.customAttrs[key]
	}
	return nil
}

func (t *CallTracker) SetRoutingRegionId(regionId string) {
	if t != nil {
		t.executedViaRegion = regionId
	}
}

func (t *CallTracker) SetExecutingRegionIds(regionIds ...string) {
	if t != nil {
		t.executedByRegions = regionIds
	}
}

func (t *CallTracker) SetLabel(key string, value string) {
	if t != nil {
		t.customLabels[key] = value
	}
}

func (t *CallTracker) SetAttribute(key string, value interface{}) {
	if t != nil {
		t.customAttrs[key] = value
	}
}

func (t *CallTracker) SetStringAttribute(key string, value string) {
	if t != nil {
		t.customAttrs[key] = value
	}
}

func (t *CallTracker) SetStringsArrayAttribute(key string, values []string) {
	if t != nil {
		t.customAttrs[key] = values
	}
}

func (t *CallTracker) SetCurrentTxTracker(txTracker *TxTracker) {
	if t != nil {
		t.txTracker = txTracker
	}
}

func (t *CallTracker) CompleteCurrentTxTracker() {
	if t != nil && t.txTracker != nil {
		if t.txTracker.err == nil {
			if len(t.successfulTxs) >= MaxRememberedTxRecordsInCallTracker {
				t.successfulTxs = slices.Delete(t.successfulTxs, 0, 1)
			}
			t.successfulTxs = append(t.successfulTxs, t.txTracker)
		} else {
			if len(t.failedTxs) >= MaxRememberedTxRecordsInCallTracker {
				t.failedTxs = slices.Delete(t.failedTxs, 0, 1)
			}
			t.failedTxs = append(t.failedTxs, t.txTracker)
		}
		t.txTracker = nil
	}
}

type ResourceChange struct {
	pre  resource.Resource
	post resource.Resource
}

type TxTracker struct {
	txId            string
	txAttempt       int
	resourceChanges []*ResourceChange
	customLabels    map[string]string
	customAttrs     map[string]interface{}
	err             error
}

func NewTxTracker(txId string) *TxTracker {
	return &TxTracker{
		txId: txId,
	}
}

func (t *TxTracker) SetExitErr(err error) {
	if t != nil {
		t.err = err
	}
}

func (t *TxTracker) BeginAttempt(attempt int) {
	if t != nil {
		t.txAttempt = attempt
		t.customLabels = map[string]string{}
		t.customAttrs = map[string]interface{}{}
		t.resourceChanges = nil
	}
}

func (t *TxTracker) SetLabel(key, value string) {
	if t != nil {
		t.customLabels[key] = value
	}
}

func (t *TxTracker) SetAttribute(key string, value interface{}) {
	if t != nil {
		t.customAttrs[key] = value
	}
}

func (t *TxTracker) AddNewResource(newRes resource.Resource) {
	if t != nil {
		t.resourceChanges = append(t.resourceChanges, &ResourceChange{
			post: newRes,
		})
	}
}

func (t *TxTracker) AddUpdatedResource(oldRes resource.Resource, newRes resource.Resource) {
	if t != nil {
		t.resourceChanges = append(t.resourceChanges, &ResourceChange{
			pre:  oldRes,
			post: newRes,
		})
	}
}

func (t *TxTracker) AddDeletedResource(delRes resource.Resource) {
	if t != nil {
		t.resourceChanges = append(t.resourceChanges, &ResourceChange{
			pre: delRes,
		})
	}
}

func (t *TxTracker) GetLabel(key string) string {
	if t != nil {
		return t.customLabels[key]
	}
	return ""
}

func (t *TxTracker) GetAttribute(key string) interface{} {
	if t != nil {
		return t.customAttrs[key]
	}
	return nil
}

func (t *TxTracker) GetResourceChanges() []*ResourceChange {
	if t != nil {
		return t.resourceChanges
	}
	return nil
}

func (t *TxTracker) GetCurrentTxId() string {
	if t != nil {
		return t.txId
	}
	return ""
}

func (t *TxTracker) GetCurrentTxAttempt() int {
	if t != nil {
		return t.txAttempt
	}
	return 0
}

func (c *ResourceChange) GetName() resource.Name {
	if c.post != nil {
		return c.post.GetRawName()
	}
	return c.pre.GetRawName()
}

func (c *ResourceChange) IsCreate() bool {
	return c.post != nil && c.pre == nil
}

func (c *ResourceChange) IsUpdate() bool {
	return c.post != nil && c.pre != nil
}

func (c *ResourceChange) IsDelete() bool {
	return c.post == nil && c.pre != nil
}

func (c *ResourceChange) GetPre() resource.Resource {
	return c.pre
}

func (c *ResourceChange) GetPost() resource.Resource {
	return c.post
}

func (c *ResourceChange) SetPre(pre resource.Resource) {
	c.pre = pre
}

func (c *ResourceChange) SetPost(post resource.Resource) {
	c.post = post
}

func extractCallerIp(clientHeader metadata.MD) string {
	// This is our internal header where original client IP must be stored
	// TODO: This is not protected from forging by Goten, implementation
	// must verify that this can be trusted. We do this in EdgeLQ in
	// authenticator
	// First original address is obtained by FIRST Goten-based service
	// from StdRealIpHeader and whenever there is a routing, we
	// store it in GotenOriginalIpAddress.
	// Therefore, in StdRealIpHeader we always take the rightmost
	// value, containing definitely trusted value.
	originalIpVals := clientHeader.Get(GotenOriginalIpAddress)
	if len(originalIpVals) > 0 && originalIpVals[0] != "" {
		return originalIpVals[0]
	}

	// TODO: We assume that there is trusted proxy (Ingress), and current
	// service cannot be connected from the internet.
	// If this service could be connected from internet directly, then
	// we should use st.RemoveAddr directly!
	realIpVals := clientHeader.Get(StdRealIpHeader)
	if len(realIpVals) > 0 {
		splitVals := strings.Split(realIpVals[len(realIpVals)-1], ",")
		if len(splitVals) > 0 {
			return strings.Join(strings.Fields(splitVals[len(splitVals)-1]), "")
		}
	}

	// try x-forwarded-for, BUT assume that very last IP address belongs to proxy...
	// TODO: Not most elegant... this is covering just as many ingresses as possible.
	forwardedIps := clientHeader.Get(StdForwardedForHeader)
	if len(forwardedIps) > 0 {
		addresses := make([]string, 0)
		for _, v := range forwardedIps {
			addresses = append(addresses, strings.Split(v, ",")...)
		}
		if len(addresses) == 1 {
			return strings.Join(strings.Fields(addresses[0]), "")
		} else if len(addresses) > 1 {
			return strings.Join(strings.Fields(addresses[len(addresses)-2]), "")
		}
	}
	return UnknownIPAddress
}
