// Code generated by protoc-gen-goten-object
// File: goten/types/memo.proto
// DO NOT EDIT!!!

package memo

// proto imports
import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// make sure we're using proto imports
var (
	_ = &timestamppb.Timestamp{}
)

type MemoFieldPathBuilder struct{}

func NewMemoFieldPathBuilder() MemoFieldPathBuilder {
	return MemoFieldPathBuilder{}
}
func (MemoFieldPathBuilder) CreateTime() MemoPathSelectorCreateTime {
	return MemoPathSelectorCreateTime{}
}
func (MemoFieldPathBuilder) UpdateTime() MemoPathSelectorUpdateTime {
	return MemoPathSelectorUpdateTime{}
}
func (MemoFieldPathBuilder) Message() MemoPathSelectorMessage {
	return MemoPathSelectorMessage{}
}
func (MemoFieldPathBuilder) CreatedBy() MemoPathSelectorCreatedBy {
	return MemoPathSelectorCreatedBy{}
}

type MemoPathSelectorCreateTime struct{}

func (MemoPathSelectorCreateTime) FieldPath() *Memo_FieldTerminalPath {
	return &Memo_FieldTerminalPath{selector: Memo_FieldPathSelectorCreateTime}
}

func (s MemoPathSelectorCreateTime) WithValue(value *timestamppb.Timestamp) *Memo_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Memo_FieldTerminalPathValue)
}

func (s MemoPathSelectorCreateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Memo_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Memo_FieldTerminalPathArrayOfValues)
}

type MemoPathSelectorUpdateTime struct{}

func (MemoPathSelectorUpdateTime) FieldPath() *Memo_FieldTerminalPath {
	return &Memo_FieldTerminalPath{selector: Memo_FieldPathSelectorUpdateTime}
}

func (s MemoPathSelectorUpdateTime) WithValue(value *timestamppb.Timestamp) *Memo_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Memo_FieldTerminalPathValue)
}

func (s MemoPathSelectorUpdateTime) WithArrayOfValues(values []*timestamppb.Timestamp) *Memo_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Memo_FieldTerminalPathArrayOfValues)
}

type MemoPathSelectorMessage struct{}

func (MemoPathSelectorMessage) FieldPath() *Memo_FieldTerminalPath {
	return &Memo_FieldTerminalPath{selector: Memo_FieldPathSelectorMessage}
}

func (s MemoPathSelectorMessage) WithValue(value string) *Memo_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Memo_FieldTerminalPathValue)
}

func (s MemoPathSelectorMessage) WithArrayOfValues(values []string) *Memo_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Memo_FieldTerminalPathArrayOfValues)
}

type MemoPathSelectorCreatedBy struct{}

func (MemoPathSelectorCreatedBy) FieldPath() *Memo_FieldTerminalPath {
	return &Memo_FieldTerminalPath{selector: Memo_FieldPathSelectorCreatedBy}
}

func (s MemoPathSelectorCreatedBy) WithValue(value string) *Memo_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*Memo_FieldTerminalPathValue)
}

func (s MemoPathSelectorCreatedBy) WithArrayOfValues(values []string) *Memo_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*Memo_FieldTerminalPathArrayOfValues)
}
