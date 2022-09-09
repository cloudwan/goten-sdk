package resource

import (
	"context"
	"reflect"
)

const (
	InConditionsMaxArgs = 10
)

func IterateListQuery(ctx context.Context, query ListQuery, access Access, fc func(results ResourceList) error) error {
	const defaultPageSize = 100

	if query.GetPager().GetLimit() == 0 {
		query.GetPager().SetLimit(defaultPageSize)
	}
	for {
		qrs, err := access.Query(ctx, query)
		if err != nil {
			return err
		}
		results := qrs.GetResults()
		if results.Length() == 0 {
			return nil
		}
		if err = fc(results); err != nil {
			return err
		}
		nextPageCursor := qrs.GetNextPageCursor()
		if nextPageCursor == nil || reflect.ValueOf(nextPageCursor).IsNil() {
			return nil
		}
		query.GetPager().SetCursor(nextPageCursor)
	}
}
