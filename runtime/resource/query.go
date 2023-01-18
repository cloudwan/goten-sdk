package resource

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cloudwan/goten-sdk/runtime/api/watch_type"
	"github.com/cloudwan/goten-sdk/runtime/object"
	"github.com/cloudwan/goten-sdk/runtime/resource/query_result"
)

const (
	DefaultMaximumQueryChunkSize = 100
)

type Query interface {
	fmt.Stringer
	GotenQuery()
	GetResourceDescriptor() Descriptor
}

type GetQuery interface {
	Query
	GetReference() Reference
	GetFieldMask() object.FieldMask
	SetReference(Reference)
	SetFieldMask(object.FieldMask)
}

type ListQuery interface {
	Query
	GetFilter() Filter
	GetPager() PagerQuery
	GetFieldMask() object.FieldMask
	GetWithPagingInfo() bool
	SetFilter(Filter)
	SetPager(PagerQuery)
	SetFieldMask(object.FieldMask)
	SetWithPagingInfo(bool)
}

type WatchQuery interface {
	ListQuery
	GetWatchType() watch_type.WatchType
	GetMaximumChunkSize() int
	GetResumeToken() string
	GetStartingTime() *timestamppb.Timestamp
	SetWatchType(watch_type.WatchType)
	SetMaximumChunkSize(int)
	SetResumeToken(string)
	SetStartingTime(*timestamppb.Timestamp)
}

type SearchQuery interface {
	ListQuery
	GetPhrase() string
	SetPhrase(string)
}

type QueryResultSnapshot interface {
	GetResults() ResourceList
	GetNextPageCursor() Cursor
	GetPrevPageCursor() Cursor
	GetPagingInfo() (totalCount, offset int32)
	SetResults(ResourceList)
	SetCursors(Cursor, Cursor)
	SetPagingInfo(totalCount, offset int32)
}

type QueryResultChange interface {
	GetResults() ResourceChangeList
	GetNextPageCursor() Cursor
	GetPrevPageCursor() Cursor
	GetIsCurrent() bool
	GetIsHardReset() bool
	GetIsSoftReset() bool
	GetSnapshotSize() int64
	GetResumeToken() string
	SetResults(ResourceChangeList)
	SetCursors(Cursor, Cursor)
	SetIsCurrent()
	SetIsHardReset()
	SetIsSoftReset()
	SetSnapshotSize(int64)
	SetResumeToken(string)
}

func MakeSQLGetString(query GetQuery) string {
	return fmt.Sprintf("SELECT %s FROM %s WHERE name = %s;",
		fieldMaskAsSQLParam(query.GetFieldMask()),
		query.GetResourceDescriptor().GetResourceTypeName().JSONPlural(),
		query.GetReference())
}

func MakeSQLListString(query ListQuery) string {
	return fmt.Sprintf("SELECT %s%s FROM %s%s%s;",
		fieldMaskAsSQLParam(query.GetFieldMask()),
		withPagingInfoAsSQLParam(query.GetWithPagingInfo()),
		query.GetResourceDescriptor().GetResourceTypeName().JSONPlural(),
		maybeAppendFilterAsSQLLikeParam(query.GetFilter()),
		maybeAppendPagerAsSQLLikeParam(query.GetPager()))
}

func MakeSQLSearchString(query SearchQuery) string {
	return fmt.Sprintf("SELECT %s%s WITH PHRASE (%s) FROM %s%s%s;",
		fieldMaskAsSQLParam(query.GetFieldMask()),
		withPagingInfoAsSQLParam(true),
		query.GetPhrase(),
		query.GetResourceDescriptor().GetResourceTypeName().JSONPlural(),
		maybeAppendFilterAsSQLLikeParam(query.GetFilter()),
		maybeAppendPagerAsSQLLikeParam(query.GetPager()))
}

func MakeSQLWatchString(query WatchQuery) string {
	return fmt.Sprintf("%s-WATCH %s FROM %s%s%s%s%s;",
		query.GetWatchType().String(),
		fieldMaskAsSQLParam(query.GetFieldMask()),
		query.GetResourceDescriptor().GetResourceTypeName().JSONPlural(),
		maybeAppendFilterAsSQLLikeParam(query.GetFilter()),
		maybeAppendPagerAsSQLLikeParam(query.GetPager()),
		maybeAppendMaxChunkSizeAsSQLLikeParam(query.GetMaximumChunkSize()),
		maybeAppendResumeTokenAsSQLLikeParam(query.GetResumeToken()))
}

func MarshalQueryResultSnapshot(qres QueryResultSnapshot) ([]byte, error) {
	if qres == nil {
		return nil, nil
	}
	totalCount, offset := qres.GetPagingInfo()
	anyResult := &query_result.QueryResult{
		TotalCount: totalCount,
		Offset:     offset,
	}
	if err := marshalResourceResults(qres.GetResults(), anyResult); err != nil {
		return nil, err
	}
	if err := marshalNextAndPrevPageCursors(qres.GetNextPageCursor(), qres.GetPrevPageCursor(), anyResult); err != nil {
		return nil, err
	}
	return proto.Marshal(anyResult)
}

func MarshalQueryResultChange(qres QueryResultChange) ([]byte, error) {
	if qres == nil {
		return nil, nil
	}
	anyResult := &query_result.QueryResult{
		IsCurrent:    qres.GetIsCurrent(),
		IsSoftReset:  qres.GetIsSoftReset(),
		IsHardReset:  qres.GetIsHardReset(),
		SnapshotSize: qres.GetSnapshotSize(),
		ResumeToken:  qres.GetResumeToken(),
	}
	if err := marshalResourceChangeResults(qres.GetResults(), anyResult); err != nil {
		return nil, err
	}
	if err := marshalNextAndPrevPageCursors(qres.GetNextPageCursor(), qres.GetPrevPageCursor(), anyResult); err != nil {
		return nil, err
	}
	return proto.Marshal(anyResult)
}

func UnmarshalQueryResultSnapshot(qres QueryResultSnapshot, descriptor Descriptor, data []byte) error {
	anyResult := &query_result.QueryResult{}
	if err := proto.Unmarshal(data, anyResult); err != nil {
		return err
	}
	if results, err := unmarshalResourceResults(descriptor, anyResult); err != nil {
		return err
	} else {
		qres.SetResults(results)
	}
	if next, prev, err := unmarshalNextAndPrevPageCursors(descriptor, anyResult); err != nil {
		return err
	} else {
		qres.SetCursors(next, prev)
	}
	qres.SetPagingInfo(anyResult.GetTotalCount(), anyResult.GetOffset())
	return nil
}

func UnmarshalQueryResultChange(qres QueryResultChange, descriptor Descriptor, data []byte) error {
	anyResult := &query_result.QueryResult{}
	if err := proto.Unmarshal(data, anyResult); err != nil {
		return err
	}
	if results, err := unmarshalResourceChangeResults(descriptor, anyResult); err != nil {
		return err
	} else {
		qres.SetResults(results)
	}
	if next, prev, err := unmarshalNextAndPrevPageCursors(descriptor, anyResult); err != nil {
		return err
	} else {
		qres.SetCursors(next, prev)
	}
	if anyResult.GetIsCurrent() {
		qres.SetIsCurrent()
	}
	if anyResult.GetIsSoftReset() {
		qres.SetIsSoftReset()
	}
	if anyResult.GetIsHardReset() {
		qres.SetIsHardReset()
	}
	qres.SetSnapshotSize(anyResult.GetSnapshotSize())
	qres.SetResumeToken(anyResult.GetResumeToken())
	return nil
}

func withPagingInfoAsSQLParam(withPagingInfo bool) string {
	if withPagingInfo {
		return " WITH COUNT"
	}
	return ""
}

func fieldMaskAsSQLParam(mask object.FieldMask) string {
	if mask == nil || reflect.ValueOf(mask).IsNil() {
		return "*"
	} else {
		return mask.String()
	}
}

func maybeAppendFilterAsSQLLikeParam(filter Filter) string {
	if filter == nil || reflect.ValueOf(filter).IsNil() {
		return ""
	} else {
		return " WHERE " + filter.String()
	}
}

func maybeAppendPagerAsSQLLikeParam(pager PagerQuery) string {
	res := ""
	if pager != nil && !reflect.ValueOf(pager).IsNil() {
		orderBy := pager.GetOrderBy()
		cursor := pager.GetCursor()
		if orderBy != nil && !reflect.ValueOf(orderBy).IsNil() {
			res += " ORDER BY " + orderBy.String()
		}
		if cursor != nil && !reflect.ValueOf(cursor).IsNil() {
			res += " WITH CURSOR " + cursor.String()
		}
		if pager.GetLimit() != 0 {
			res += fmt.Sprintf(" LIMIT %d", pager.GetLimit())
		}
	}
	return res
}

func maybeAppendMaxChunkSizeAsSQLLikeParam(chunkSize int) string {
	if chunkSize > 0 {
		return fmt.Sprintf(" WITH MAX-CHUNK-SIZE = %d", chunkSize)
	}
	return ""
}

func maybeAppendResumeTokenAsSQLLikeParam(token string) string {
	if token == "" {
		return fmt.Sprintf(" WITH RESUME-TOKEN = %s", token)
	}
	return ""
}

func marshalResourceResults(results ResourceList, output *query_result.QueryResult) error {
	output.Results = make([]*anypb.Any, 0, results.Length())
	for i := 0; i < results.Length(); i++ {
		result := results.At(i)
		if anyObj, err := anypb.New(result); err != nil {
			return err
		} else {
			output.Results = append(output.Results, anyObj)
		}
	}
	return nil
}

func marshalResourceChangeResults(results ResourceChangeList, output *query_result.QueryResult) error {
	output.Results = make([]*anypb.Any, 0, results.Length())
	for i := 0; i < results.Length(); i++ {
		result := results.At(i)
		if anyObj, err := anypb.New(result); err != nil {
			return err
		} else {
			output.Results = append(output.Results, anyObj)
		}
	}
	return nil
}

func marshalNextAndPrevPageCursors(nextPageCursor, prevPageCursor Cursor, output *query_result.QueryResult) error {
	if nextPageCursor != nil {
		nValue, err := nextPageCursor.ProtoString()
		if err != nil {
			return err
		}
		output.NextCursor = []byte(nValue)
	}
	if prevPageCursor != nil {
		pValue, err := prevPageCursor.ProtoString()
		if err != nil {
			return err
		}
		output.PrevCursor = []byte(pValue)
	}
	return nil
}

func unmarshalResourceResults(descriptor Descriptor, input *query_result.QueryResult) (ResourceList, error) {
	resources := descriptor.NewResourceList(0, len(input.Results))
	for _, result := range input.Results {
		res := descriptor.NewResource()
		if err := anypb.UnmarshalTo(result, res, proto.UnmarshalOptions{}); err != nil {
			return nil, err
		}
		resources = resources.Append(res)
	}
	return resources, nil
}

func unmarshalResourceChangeResults(descriptor Descriptor, input *query_result.QueryResult) (ResourceChangeList, error) {
	changes := descriptor.NewResourceChangeList(0, len(input.Results))
	for _, result := range input.Results {
		resChange := descriptor.NewResourceChange()
		if err := anypb.UnmarshalTo(result, resChange, proto.UnmarshalOptions{}); err != nil {
			return nil, err
		}
		changes = changes.Append(resChange)
	}
	return changes, nil
}

func unmarshalNextAndPrevPageCursors(descriptor Descriptor, input *query_result.QueryResult) (Cursor, Cursor, error) {
	var nextPageCursor Cursor
	var prevPageCursor Cursor
	if len(input.NextCursor) > 0 {
		nextPageCursor = descriptor.NewResourceCursor()
		if err := nextPageCursor.ParseProtoString(string(input.NextCursor)); err != nil {
			return nil, nil, err
		}
	}
	if len(input.PrevCursor) > 0 {
		prevPageCursor = descriptor.NewResourceCursor()
		if err := prevPageCursor.ParseProtoString(string(input.PrevCursor)); err != nil {
			return nil, nil, err
		}
	}
	return nextPageCursor, prevPageCursor, nil
}
