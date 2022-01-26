package object

import (
	"reflect"

	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// special reflection around any field mask.
// In proto, each field mask is google.protobuf.FieldMask
// In Go, we have our own custom type. We must provide to every
// our custom FieldMask preflect.Message which will map between
// google.protobuf.FieldMask and our custom <Object>_FieldMask.

const (
	PathsFieldName = "paths"
)

type fieldMaskType struct {
	mask            FieldMask
	protoDescriptor preflect.MessageDescriptor
}

func (maskType *fieldMaskType) New() preflect.Message {
	newMsg := reflect.New(reflect.TypeOf(maskType.mask).Elem()).Interface().(preflect.ProtoMessage)
	return newMsg.ProtoReflect()
}

func (maskType *fieldMaskType) Zero() preflect.Message {
	zeroMsg := reflect.Zero(reflect.TypeOf(maskType.mask)).Interface().(preflect.ProtoMessage)
	return zeroMsg.ProtoReflect()
}

func (maskType *fieldMaskType) Descriptor() preflect.MessageDescriptor {
	return maskType.protoDescriptor
}

type fieldPathList struct {
	mask       FieldMask
	pathParser func(raw string) (FieldPath, error)
}

func (l *fieldPathList) Len() int {
	if !l.IsValid() {
		return 0
	}
	return len(l.mask.GetRawPaths())
}

func (l *fieldPathList) Get(i int) preflect.Value {
	path := l.mask.GetRawPaths()[i]
	return preflect.ValueOfString(path.String())
}

func (l *fieldPathList) Set(i int, v preflect.Value) {
	currentPaths := l.mask.GetRawPaths()
	path, err := l.pathParser(v.String())
	if err != nil {
		panic(err)
	}
	currentPaths[i] = path
	proto.Reset(l.mask)
	for _, path := range currentPaths {
		l.mask.AppendRawPath(path)
	}
}

func (l *fieldPathList) Append(v preflect.Value) {
	path, err := l.pathParser(v.String())
	if err != nil {
		panic(err)
	}
	l.mask.AppendRawPath(path)
}

func (l *fieldPathList) AppendMutable() preflect.Value {
	// cant, have fun
	return preflect.ValueOfString("")
}

func (l *fieldPathList) Truncate(n int) {
	currentPaths := l.mask.GetRawPaths()
	if n < len(currentPaths) {
		proto.Reset(l.mask)
		for _, path := range currentPaths[:n] {
			l.mask.AppendRawPath(path)
		}
	}
}

func (l *fieldPathList) NewElement() preflect.Value {
	// cant, have fun
	return preflect.ValueOfString("")
}

func (l *fieldPathList) IsValid() bool {
	return !reflect.ValueOf(l.mask).IsNil()
}

type fieldMaskReflection struct {
	PathParser     func(raw string) (FieldPath, error)
	Mask           FieldMask
	BaseReflection preflect.Message
}

func MakeFieldMaskReflection(mask FieldMask, pathParser func(raw string) (FieldPath, error)) preflect.Message {
	return &fieldMaskReflection{
		PathParser:     pathParser,
		Mask:           mask,
		BaseReflection: (*fieldmaskpb.FieldMask)(nil).ProtoReflect(),
	}
}

func (reflection *fieldMaskReflection) Descriptor() preflect.MessageDescriptor {
	return reflection.BaseReflection.Descriptor()
}

func (reflection *fieldMaskReflection) Type() preflect.MessageType {
	return &fieldMaskType{
		mask:            reflection.Mask,
		protoDescriptor: reflection.Descriptor(),
	}
}

func (reflection *fieldMaskReflection) New() preflect.Message {
	newMsg := reflect.New(reflect.TypeOf(reflection.Mask).Elem()).Interface().(preflect.ProtoMessage)
	return newMsg.ProtoReflect()
}

func (reflection *fieldMaskReflection) Interface() preflect.ProtoMessage {
	return reflection.Mask
}

func (reflection *fieldMaskReflection) Range(f func(preflect.FieldDescriptor, preflect.Value) bool) {
	if !reflection.IsValid() {
		return
	}
	fields := reflection.Descriptor().Fields()
	protoMask := reflection.Mask.ToProtoFieldMask()
	protoReflect := protoMask.ProtoReflect()
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		if protoReflect.Has(fd) {
			v := protoReflect.Get(fd)
			f(fd, v)
		}
	}
	if err := reflection.Mask.FromProtoFieldMask(protoMask); err != nil {
		panic(err)
	}
}

func (reflection *fieldMaskReflection) Has(fd preflect.FieldDescriptor) bool {
	if !reflection.IsValid() {
		return false
	}
	return reflection.Mask.ToProtoFieldMask().ProtoReflect().Has(fd)
}

func (reflection *fieldMaskReflection) Clear(fd preflect.FieldDescriptor) {
	protoMask := reflection.Mask.ToProtoFieldMask()
	protoMask.ProtoReflect().Clear(fd)
	if err := reflection.Mask.FromProtoFieldMask(protoMask); err != nil {
		panic(err)
	}
}

func (reflection *fieldMaskReflection) Get(fd preflect.FieldDescriptor) preflect.Value {
	protoMask := reflection.Mask.ToProtoFieldMask()
	return protoMask.ProtoReflect().Get(fd)
}

func (reflection *fieldMaskReflection) Set(fd preflect.FieldDescriptor, v preflect.Value) {
	protoMask := reflection.Mask.ToProtoFieldMask()
	protoMask.ProtoReflect().Set(fd, v)
	if err := reflection.Mask.FromProtoFieldMask(protoMask); err != nil {
		panic(err)
	}
}

func (reflection *fieldMaskReflection) Mutable(fd preflect.FieldDescriptor) preflect.Value {
	if fd.Name() == PathsFieldName {
		return preflect.ValueOfList(&fieldPathList{mask: reflection.Mask, pathParser: reflection.PathParser})
	} else {
		// this is WKT FieldMask, so actually should not happen
		protoMask := reflection.Mask.ToProtoFieldMask()
		return protoMask.ProtoReflect().Mutable(fd)
	}
}

func (reflection *fieldMaskReflection) NewField(fd preflect.FieldDescriptor) preflect.Value {
	reflection.Clear(fd)
	if fd.Name() == PathsFieldName {
		return preflect.ValueOfList(&fieldPathList{mask: reflection.Mask, pathParser: reflection.PathParser})
	} else {
		// this is WKT FieldMask, so actually should not happen
		protoMask := reflection.Mask.ToProtoFieldMask()
		return protoMask.ProtoReflect().Get(fd)
	}
}

func (reflection *fieldMaskReflection) WhichOneof(preflect.OneofDescriptor) preflect.FieldDescriptor {
	return nil
}

func (reflection *fieldMaskReflection) GetUnknown() preflect.RawFields {
	return nil
}

func (reflection *fieldMaskReflection) SetUnknown(preflect.RawFields) {}

func (reflection *fieldMaskReflection) IsValid() bool {
	return !reflect.ValueOf(reflection.Mask).IsNil()
}

func (reflection *fieldMaskReflection) ProtoMethods() *protoiface.Methods {
	methods := reflection.BaseReflection.ProtoMethods()

	return &protoiface.Methods{
		Flags: methods.Flags,
		Marshal: func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
			gotenMask := input.Message.Interface().(FieldMask)
			if reflect.ValueOf(gotenMask).IsNil() {
				return protoiface.MarshalOutput{Buf: input.Buf}, nil
			}
			opts := proto.MarshalOptions{
				Deterministic: (input.Flags & protoiface.MarshalDeterministic) != 0,
				UseCachedSize: (input.Flags & protoiface.MarshalUseCachedSize) != 0,
			}
			protoMask := gotenMask.ToProtoFieldMask()
			data, err := opts.Marshal(protoMask)
			if err != nil {
				return protoiface.MarshalOutput{Buf: input.Buf}, err
			}
			return protoiface.MarshalOutput{Buf: append(input.Buf, data...)}, nil
		},
		Unmarshal: func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
			gotenMask := input.Message.Interface().(FieldMask)
			googleMask := &fieldmaskpb.FieldMask{}
			if err := proto.Unmarshal(input.Buf, googleMask); err != nil {
				return protoiface.UnmarshalOutput{}, err
			}
			if err := gotenMask.FromProtoFieldMask(googleMask); err != nil {
				return protoiface.UnmarshalOutput{}, err
			}
			return protoiface.UnmarshalOutput{}, nil
		},
		Merge: func(input protoiface.MergeInput) protoiface.MergeOutput {
			srcMask := input.Source.Interface().(FieldMask)
			dstMask := input.Destination.Interface().(FieldMask)
			if reflect.ValueOf(srcMask).IsNil() {
				return protoiface.MergeOutput{}
			}
			for _, path := range srcMask.GetRawPaths() {
				dstMask.AppendRawPath(path)
			}
			return protoiface.MergeOutput{
				Flags: protoiface.MergeComplete,
			}
		},
		Size: func(input protoiface.SizeInput) protoiface.SizeOutput {
			gotenMask := input.Message.Interface().(FieldMask)
			if reflect.ValueOf(gotenMask).IsNil() {
				return protoiface.SizeOutput{Size: 0}
			}
			protoMask := gotenMask.ToProtoFieldMask()
			return protoiface.SizeOutput{Size: proto.Size(protoMask)}
		},
	}
}
