package resource

import (
	"context"
)

func GetCount(ctx context.Context, desc Descriptor, access Access, filter Filter) (int32, error) {
	q := MakeCountQuery(desc, filter)

	qr, err := access.Query(ctx, q)
	if err != nil {
		return -1, err
	}
	count, _ := qr.GetPagingInfo()
	return count, nil
}
