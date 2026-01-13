package resource

import (
	"bytes"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cloudwan/goten-sdk/runtime/utils"
)

// IsSemanticallyEmpty returns true if msg contains no semantically-meaningful values.
// Treats nil, {}, {Nested:{}}, {Nested:{Scalar:default}} as empty.
//
// Rules:
//   - Scalars: meaningful if Range() visits them (non-default for proto3)
//   - Messages: meaningful only if they contain meaningful values (recursive)
//   - Lists: meaningful if len > 1, OR len == 1 with non-empty message element
//   - Maps: meaningful if any entry has a non-empty value
func IsSemanticallyEmpty(msg proto.Message) bool {
	if utils.IsNil(msg) {
		return true
	}

	m := msg.ProtoReflect()
	if !m.IsValid() {
		return true
	}

	empty := true
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		switch {
		case fd.IsList():
			list := v.List()
			// Special case: single empty message element treated as empty list
			if list.Len() == 1 && (fd.Kind() == protoreflect.MessageKind || fd.Kind() == protoreflect.GroupKind) {
				if IsSemanticallyEmpty(list.Get(0).Message().Interface()) {
					return true // continue, doesn't contribute to non-emptiness
				}
			}
			// Any other non-empty list is meaningful
			empty = false
			return false

		case fd.IsMap():
			// Map is meaningful only if it has entries with non-empty values
			mp := v.Map()
			mp.Range(func(k protoreflect.MapKey, mv protoreflect.Value) bool {
				if isEffectiveMapValue(fd.MapValue(), mv) {
					empty = false
					return false
				}
				return true
			})
			if !empty {
				return false
			}

		case fd.Kind() == protoreflect.MessageKind || fd.Kind() == protoreflect.GroupKind:
			// Message field: meaningful only if it has meaningful content
			if !IsSemanticallyEmpty(v.Message().Interface()) {
				empty = false
				return false
			}

		default:
			// Scalar/enum/bytes visited by Range() means non-default value
			empty = false
			return false
		}
		return true
	})

	return empty
}

// SemanticProtoEqual compares two protos treating empty submessages as absent.
// Returns true if both are semantically empty, or both have identical meaningful content.
func SemanticProtoEqual(a, b proto.Message) bool {
	aEmpty := IsSemanticallyEmpty(a)
	bEmpty := IsSemanticallyEmpty(b)
	if aEmpty || bEmpty {
		return aEmpty && bEmpty
	}
	// Both non-empty: do deep semantic comparison
	return semanticEqual(a.ProtoReflect(), b.ProtoReflect())
}

func semanticEqual(a, b protoreflect.Message) bool {
	if !a.IsValid() || !b.IsValid() {
		return !a.IsValid() && !b.IsValid()
	}
	if a.Descriptor() != b.Descriptor() {
		return false
	}

	fields := a.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)

		ha, va := effectiveGet(a, fd)
		hb, vb := effectiveGet(b, fd)

		if !ha && !hb {
			continue
		}
		if ha != hb {
			return false
		}
		if !equalField(fd, va, vb) {
			return false
		}
	}
	return true
}

// effectiveGet returns whether a field is "effectively present" and its value.
// Treats semantically-empty messages, single-empty-element lists, and
// maps with only empty values as absent.
func effectiveGet(m protoreflect.Message, fd protoreflect.FieldDescriptor) (bool, protoreflect.Value) {
	if !m.Has(fd) {
		return false, protoreflect.Value{}
	}
	v := m.Get(fd)

	if fd.IsList() {
		list := v.List()
		if list.Len() == 0 {
			return false, protoreflect.Value{}
		}
		// Single empty message element treated as absent
		if list.Len() == 1 && (fd.Kind() == protoreflect.MessageKind || fd.Kind() == protoreflect.GroupKind) {
			if IsSemanticallyEmpty(list.Get(0).Message().Interface()) {
				return false, protoreflect.Value{}
			}
		}
		return true, v
	}

	if fd.IsMap() {
		// Count only entries with effective values
		effectiveCount := 0
		v.Map().Range(func(k protoreflect.MapKey, mv protoreflect.Value) bool {
			if isEffectiveMapValue(fd.MapValue(), mv) {
				effectiveCount++
			}
			return true
		})
		if effectiveCount == 0 {
			return false, protoreflect.Value{}
		}
		return true, v
	}

	if fd.Kind() == protoreflect.MessageKind || fd.Kind() == protoreflect.GroupKind {
		if IsSemanticallyEmpty(v.Message().Interface()) {
			return false, protoreflect.Value{}
		}
	}

	return true, v
}

// isEffectiveMapValue returns true if a map value is semantically meaningful.
func isEffectiveMapValue(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
	if !v.IsValid() {
		return false
	}
	if fd.Kind() == protoreflect.MessageKind || fd.Kind() == protoreflect.GroupKind {
		return !IsSemanticallyEmpty(v.Message().Interface())
	}
	// Scalar map values are always effective if present
	return true
}

func equalField(fd protoreflect.FieldDescriptor, a, b protoreflect.Value) bool {
	switch {
	case fd.IsList():
		return equalList(fd, a.List(), b.List())
	case fd.IsMap():
		return equalMap(fd, a.Map(), b.Map())
	default:
		return equalValue(fd, a, b)
	}
}

func equalList(fd protoreflect.FieldDescriptor, la, lb protoreflect.List) bool {
	if la.Len() != lb.Len() {
		return false
	}
	for i := 0; i < la.Len(); i++ {
		if !equalValue(fd, la.Get(i), lb.Get(i)) {
			return false
		}
	}
	return true
}

func equalMap(fd protoreflect.FieldDescriptor, ma, mb protoreflect.Map) bool {
	valueDesc := fd.MapValue()

	// Count effective entries
	effectiveLenA := 0
	ma.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
		if isEffectiveMapValue(valueDesc, v) {
			effectiveLenA++
		}
		return true
	})
	effectiveLenB := 0
	mb.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
		if isEffectiveMapValue(valueDesc, v) {
			effectiveLenB++
		}
		return true
	})
	if effectiveLenA != effectiveLenB {
		return false
	}

	// Compare effective entries
	ok := true
	ma.Range(func(k protoreflect.MapKey, va protoreflect.Value) bool {
		if !isEffectiveMapValue(valueDesc, va) {
			return true // skip empty entries
		}
		vb := mb.Get(k)
		if !isEffectiveMapValue(valueDesc, vb) {
			ok = false
			return false
		}
		if !equalValue(valueDesc, va, vb) {
			ok = false
			return false
		}
		return true
	})
	return ok
}

func equalValue(fd protoreflect.FieldDescriptor, a, b protoreflect.Value) bool {
	switch fd.Kind() {
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return semanticEqual(a.Message(), b.Message())
	case protoreflect.BoolKind:
		return a.Bool() == b.Bool()
	case protoreflect.EnumKind:
		return a.Enum() == b.Enum()
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return a.Int() == b.Int()
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return a.Uint() == b.Uint()
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return a.Float() == b.Float()
	case protoreflect.StringKind:
		return a.String() == b.String()
	case protoreflect.BytesKind:
		return bytes.Equal(a.Bytes(), b.Bytes())
	default:
		return a.Interface() == b.Interface()
	}
}
