package resource

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

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
	SetFilter(Filter)
	SetPager(PagerQuery)
	SetFieldMask(object.FieldMask)
}

type WatchQuery interface {
	ListQuery
	GetWatchType() watch_type.WatchType
	GetMaximumChunkSize() int
	GetResumeToken() string
	SetWatchType(watch_type.WatchType)
	SetMaximumChunkSize(int)
	SetResumeToken(string)
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
	SetResults(ResourceList)
	SetCursors(Cursor, Cursor)
}

type SearchQueryResultSnapshot interface {
	QueryResultSnapshot
	GetPagingInfo() (totalCount, offset int32)
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

func MakeSQLGetString(mask object.FieldMask, name Name) string {
	res := "SELECT "
	if mask == nil || reflect.ValueOf(mask).IsNil() {
		res += "*"
	} else {
		res += mask.String()
	}
	res += fmt.Sprintf(" FROM %s", name.GetResourceDescriptor().GetResourceTypeName().JSONPlural())
	res += fmt.Sprintf(" WHERE name = %s", name)
	res += ";"
	return res
}

func MakeSQLString(mask object.FieldMask, filter Filter, pager PagerQuery, searchPhrase string) string {
	res := "SELECT "
	if mask == nil || reflect.ValueOf(mask).IsNil() {
		res += "*"
	} else {
		res += mask.String()
	}
	res += fmt.Sprintf(" FROM %s", pager.GetResourceDescriptor().GetResourceTypeName().JSONPlural())
	if searchPhrase != "" {
		res += " WITH SEARCH-PHRASE " + searchPhrase
	}
	if filter != nil && !reflect.ValueOf(filter).IsNil() {
		res += " WHERE " + filter.String()
	}
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
	res += ";"
	return res
}

func MarshalQueryResultSnapshot(qres QueryResultSnapshot) ([]byte, error) {
	if qres == nil {
		return nil, nil
	}
	anyQueryResult := &query_result.QueryResult{}
	results := qres.GetResults()
	for i := 0; i < results.Length(); i++ {
		result := results.At(i)
		if anyObj, err := anypb.New(result); err != nil {
			return nil, err
		} else {
			anyQueryResult.Results = append(anyQueryResult.Results, anyObj)
		}
	}
	if qres.GetNextPageCursor() != nil {
		nValue, err := qres.GetNextPageCursor().ProtoString()
		if err != nil {
			return nil, err
		}
		anyQueryResult.NextCursor = []byte(nValue)
	}
	if qres.GetPrevPageCursor() != nil {
		pValue, err := qres.GetPrevPageCursor().ProtoString()
		if err != nil {
			return nil, err
		}
		anyQueryResult.PrevCursor = []byte(pValue)
	}
	return proto.Marshal(anyQueryResult)
}

func UnmarshalQueryResultSnapshot(qres QueryResultSnapshot, descriptor Descriptor, data []byte) error {
	anyQueryResult := &query_result.QueryResult{}
	if err := proto.Unmarshal(data, anyQueryResult); err != nil {
		return err
	}

	resources := descriptor.NewResourceList(0, len(anyQueryResult.Results))
	for _, result := range anyQueryResult.Results {
		res := descriptor.NewResource()
		if err := anypb.UnmarshalTo(result, res, proto.UnmarshalOptions{}); err != nil {
			return err
		}
		resources = resources.Append(res)
	}

	var nextPageCursor Cursor
	var prevPageCursor Cursor
	if len(anyQueryResult.NextCursor) > 0 {
		nextPageCursor = descriptor.NewResourceCursor()
		if err := nextPageCursor.ParseProtoString(string(anyQueryResult.NextCursor)); err != nil {
			return err
		}
	}
	if len(anyQueryResult.PrevCursor) > 0 {
		prevPageCursor = descriptor.NewResourceCursor()
		if err := prevPageCursor.ParseProtoString(string(anyQueryResult.PrevCursor)); err != nil {
			return err
		}
	}
	qres.SetCursors(nextPageCursor, prevPageCursor)
	qres.SetResults(resources)
	return nil
}
