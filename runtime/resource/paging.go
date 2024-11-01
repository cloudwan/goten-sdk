package resource

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cloudwan/goten-sdk/runtime/object"
)

type OrderByField interface {
	GetDirection() OrderDirection
	GetFieldPath() object.FieldPath
}

// Cursor wraps information about point in database index, direction and inclusiveness.
type Cursor interface {
	fmt.Stringer
	preflect.ProtoStringer
	IsEmpty() bool
	GetPageDirection() PageDirection
	GetInclusion() CursorInclusion
	GetValue() CursorValue
	SetPageDirection(PageDirection)
	SetInclusion(CursorInclusion)
	SetCursorValue(CursorValue)
}

// CursorValue is implementation-specific holder of value indicating position of the cursor in database index.
type CursorValue interface {
	fmt.Stringer
	GetValueType() CursorValueType
}

// SnapshotCursorValue uses resource as a point in the database index
type SnapshotCursorValue struct {
	Snapshot Resource
}

func ParseSnapshotCursorValue(descriptor Descriptor, strValue string) (*SnapshotCursorValue, error) {
	rawSnapshot, err := base64.StdEncoding.DecodeString(strValue)
	if err != nil {
		return nil, err
	}
	snapshot := descriptor.NewResource()
	if err = proto.Unmarshal(rawSnapshot, snapshot); err != nil {
		return nil, err
	}
	return &SnapshotCursorValue{Snapshot: snapshot}, nil
}

func NewSnapshotCursorValue(snapshot Resource) *SnapshotCursorValue {
	return &SnapshotCursorValue{Snapshot: snapshot}
}

func (scp *SnapshotCursorValue) GetSnapshot() Resource {
	if scp == nil {
		return nil
	}
	return scp.Snapshot
}

func (scp *SnapshotCursorValue) String() string {
	if scp == nil {
		return ""
	}
	snap, err := proto.Marshal(scp.Snapshot)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(snap)
}

func (scp *SnapshotCursorValue) GetValueType() CursorValueType {
	return SnapshotCursorValueType
}

// OffsetCursorValue uses integer offset as a way to indicate point in database for Cursor.
// Offset must be 0-indexed and point from the beginning for given Filter and OrderBy.
type OffsetCursorValue struct {
	Offset int32
}

func ParseOffsetCursorValue(strValue string) (*OffsetCursorValue, error) {
	offset, err := strconv.Atoi(strValue)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot parse %s as Offset Cursor: %s", strValue, err)
	}
	return &OffsetCursorValue{Offset: int32(offset)}, nil
}

func NewOffsetCursorValue(offset int32) *OffsetCursorValue {
	return &OffsetCursorValue{Offset: offset}
}

func (ocp *OffsetCursorValue) GetOffset() int32 {
	if ocp == nil {
		return 0
	}
	return ocp.Offset
}

func (ocp *OffsetCursorValue) MakeForwardValue(direction PageDirection, advanceBy int32) *OffsetCursorValue {
	offset := ocp.GetOffset()
	if direction.IsRight() {
		offset += advanceBy
	} else {
		offset -= advanceBy
		if offset < 0 {
			offset = 0
		}
	}
	return &OffsetCursorValue{Offset: offset}
}

func (ocp *OffsetCursorValue) MakeBackwardValue(direction PageDirection, advanceBy int32) *OffsetCursorValue {
	offset := ocp.GetOffset()
	if direction.IsRight() {
		offset -= advanceBy
		if offset < 0 {
			offset = 0
		}
	} else {
		offset += advanceBy
	}
	return &OffsetCursorValue{Offset: offset}
}

func (ocp *OffsetCursorValue) String() string {
	if ocp == nil {
		return ""
	}
	return fmt.Sprintf("%d", ocp.Offset)
}

func (ocp *OffsetCursorValue) GetValueType() CursorValueType {
	return OffsetCursorValueType
}

// CustomCursorValue uses custom-made params as a way to indicate point in database for Cursor.
// It can be used by any backend in a way it likes.
type CustomCursorValue struct {
	Params map[string]string
}

func ParseCustomCursorValue(strValue string) (*CustomCursorValue, error) {
	customParams := strings.Split(strValue, ",")
	params := make(map[string]string, len(customParams))
	for _, keyValue := range customParams {
		split := strings.Split(keyValue, ":")
		if len(split) != 2 {
			return nil, status.Errorf(codes.InvalidArgument,
				"CustomCursorValue must consist of key values separated by ':' got %s", keyValue)
		}
		keyBytes, err := base64.StdEncoding.DecodeString(split[0])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument,
				"error decoding key while parsing custom cursor value, %s is not valid base64: %s", split[0], err)
		}
		valueBytes, err := base64.StdEncoding.DecodeString(split[1])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument,
				"error decoding value while parsing custom cursor value, %s is not valid base64: %s", split[1], err)
		}
		params[string(keyBytes)] = string(valueBytes)
	}
	return &CustomCursorValue{Params: params}, nil
}

func NewCustomCursorValue() *CustomCursorValue {
	return &CustomCursorValue{Params: map[string]string{}}
}

func (occ *CustomCursorValue) String() string {
	if occ == nil {
		return ""
	}
	items := make([]string, 0, len(occ.Params))
	for key, value := range occ.Params {
		b64Key := base64.StdEncoding.EncodeToString([]byte(key))
		b64Value := base64.StdEncoding.EncodeToString([]byte(value))
		items = append(items, fmt.Sprintf("%s:%s", b64Key, b64Value))
	}
	return strings.Join(items, ",")
}

func (occ *CustomCursorValue) GetValueType() CursorValueType {
	return CustomCursorValueType
}

type CursorValueType string

const (
	UnIdentifiedCursorValue CursorValueType = ""
	SnapshotCursorValueType CursorValueType = "S"
	OffsetCursorValueType   CursorValueType = "O"
	CustomCursorValueType   CursorValueType = "C"
)

// OrderBy indicates which ordering index should be used when querying for resources.
type OrderBy interface {
	fmt.Stringer
	preflect.ProtoStringer
	GetRawFieldMask() object.FieldMask
	GetOrderByFields() []OrderByField
	SortRaw(ResourceList)
	InsertSortedRaw(ResourceList, Resource) (ResourceList, int)
	CompareRaw(Resource, Resource) int
}

type OrderDirection string

const (
	DirectionASC  OrderDirection = "asc"
	DirectionDESC OrderDirection = "desc"
)

func OrderDirectionFromString(str string) (result OrderDirection, err error) {
	switch OrderDirection(strings.TrimSpace(strings.ToLower(str))) {
	case DirectionASC:
		return DirectionASC, nil
	case DirectionDESC:
		return DirectionDESC, nil
	default:
		err = status.Errorf(codes.InvalidArgument, "invalid order by direction: '%s', allowed values are: '%s', '%s'", str, DirectionASC, DirectionASC)
		return
	}
}

func (od OrderDirection) Reverse() OrderDirection {
	if od == DirectionASC {
		return DirectionDESC
	} else {
		return DirectionASC
	}
}

type PageDirection string

const (
	PageLeft  PageDirection = "l"
	PageRight PageDirection = "r"
)

func (pd PageDirection) IsRight() bool {
	return pd != PageLeft
}

func (pd PageDirection) IsLeft() bool {
	return !pd.IsRight()
}

func (pd PageDirection) Reverse() PageDirection {
	if pd.IsRight() {
		return PageLeft
	} else {
		return PageRight
	}
}

func PageDirectionFromString(str string) (result PageDirection, err error) {
	switch PageDirection(strings.TrimSpace(strings.ToLower(str))) {
	case PageLeft:
		return PageLeft, nil
	case PageRight:
		return PageRight, nil
	default:
		err = status.Errorf(codes.InvalidArgument, "invalid page direction: '%s', allowed values are: '%s', '%s'", str, PageLeft, PageRight)
		return
	}
}

type CursorInclusion string

const (
	CursorInclusive CursorInclusion = "i"
	CursorExclusive CursorInclusion = "e"
)

func (ci CursorInclusion) IsInclusive() bool {
	return ci != CursorExclusive
}

func (ci CursorInclusion) IsExclusive() bool {
	return !ci.IsInclusive()
}

func (ci CursorInclusion) Reverse() CursorInclusion {
	if ci.IsInclusive() {
		return CursorExclusive
	} else {
		return CursorInclusive
	}
}

func CursorInclusionFromString(str string) (result CursorInclusion, err error) {
	switch CursorInclusion(strings.TrimSpace(strings.ToLower(str))) {
	case CursorInclusive:
		return CursorInclusive, nil
	case CursorExclusive:
		return CursorExclusive, nil
	default:
		err = status.Errorf(codes.InvalidArgument, "invalid cursor inclusion: '%s', allowed values are: '%s', '%s'", str, CursorInclusive, CursorExclusive)
		return
	}
}
