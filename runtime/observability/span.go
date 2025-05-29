package observability

import (
	"context"
	"sync"
	"time"
)

type spanContextKeyType struct{}

var spanContextKey spanContextKeyType

func CurrentSpanFromContext(ctx context.Context) *Span {
	span := ctx.Value(spanContextKey)
	if span != nil {
		return span.(*Span)
	}
	return nil
}

type SpanCheckpoint struct {
	name string
	time time.Time
	span *Span
}

func (c *SpanCheckpoint) GetName() string {
	return c.name
}

func (c *SpanCheckpoint) GetTime() time.Time {
	return c.time
}

func (c *SpanCheckpoint) GetAssociatedSpan() *Span {
	return c.span
}

type Span struct {
	name        string
	begin, end  time.Time
	attrs       map[string]interface{}
	checkpoints []*SpanCheckpoint
	kids        []*Span
	parent      *Span
	lock        sync.Mutex
	err         error
}

func (s *Span) GetName() string {
	if s == nil {
		return "<nil>"
	}
	return s.name
}

func (s *Span) GetBeginTime() time.Time {
	if s == nil {
		return time.Time{}
	}
	return s.begin
}

func (s *Span) GetEndTime() time.Time {
	if s == nil {
		return time.Time{}
	}
	return s.end
}

func (s *Span) GetAttributes() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.attrs
}

func (s *Span) GetSpanKids() []*Span {
	if s == nil {
		return nil
	}
	return s.kids
}

func (s *Span) GetCheckpoints() []*SpanCheckpoint {
	if s == nil {
		return nil
	}
	return s.checkpoints
}

func (s *Span) RecordCheckpoint(name string) {
	if s != nil && s.end.IsZero() {
		s.checkpoints = append(s.checkpoints, &SpanCheckpoint{name: name, time: time.Now().UTC()})
	}
}

func (s *Span) GetErr() error {
	if s == nil {
		return nil
	}
	return s.err
}

func (s *Span) End(err error) {
	if s != nil {
		s.enforceEndDown(err, time.Now().UTC())
	}
}

func (s *Span) EndWithTime(err error, t time.Time) {
	if s != nil {
		s.enforceEndDown(err, t)
	}
}

func (s *Span) AddAttribute(key string, value interface{}) *Span {
	if s != nil && s.end.IsZero() {
		s.attrs[key] = value
	}
	return s
}

func (s *Span) enforceEndDown(err error, t time.Time) {
	if !s.end.IsZero() {
		return
	}
	s.end = t
	s.err = err

	s.lock.Lock()
	for _, kid := range s.kids {
		kid.enforceEndDown(err, t)
	}
	s.lock.Unlock()
}

func StartRootSpan(ctx context.Context, name string) (context.Context, *Span) {
	newSpan := &Span{
		name:  name,
		begin: time.Now().UTC(),
		attrs: make(map[string]interface{}),
	}
	ctx = context.WithValue(ctx, spanContextKey, newSpan)
	return ctx, newSpan
}

func StartSpan(ctx context.Context, name string) (context.Context, *Span) {
	parentSpan := ctx.Value(spanContextKey)

	var newSpan *Span
	for parentSpan != nil {
		typedParentSpan := parentSpan.(*Span)
		if typedParentSpan == nil {
			return ctx, nil
		}
		typedParentSpan.lock.Lock()
		isNotCompleted := typedParentSpan.end.IsZero()
		if isNotCompleted {
			newSpan = &Span{
				name:  name,
				begin: time.Now().UTC(),
				attrs: make(map[string]interface{}),
			}
			typedParentSpan.kids = append(typedParentSpan.kids, newSpan)
			typedParentSpan.checkpoints = append(typedParentSpan.checkpoints, &SpanCheckpoint{
				name: "KidSpanStarted",
				time: newSpan.begin,
				span: newSpan,
			})
			newSpan.parent = typedParentSpan
			typedParentSpan.lock.Unlock()
			break
		} else {
			// typically, case may be like:
			// ctx, span := StartSpan(ctx, ...)
			// defer span.End(err)
			// ...
			// return ctx // ctx contains finished span
			//
			// In case like above, try next parent span, which is NOT closed...
			// Perhaps caller is incorrect.
			// If root parent is complete, we will just return nil, as whole span
			// business finished by now.
			parentSpan = typedParentSpan.parent
		}
		typedParentSpan.lock.Unlock()
	}
	if newSpan == nil {
		return ctx, nil
	}
	ctx = context.WithValue(ctx, spanContextKey, newSpan)
	return ctx, newSpan
}
