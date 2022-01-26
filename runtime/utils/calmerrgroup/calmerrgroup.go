package calmerrgroup

import (
	"context"
	"fmt"
	"runtime/debug"

	"golang.org/x/sync/errgroup"
)

// Wrapper on errgroup that suppresses panics, returning error instead
type Group struct {
	inner *errgroup.Group
	ctx   context.Context
}

func WithContext(ctx context.Context) (*Group, context.Context) {
	inner, gctx := errgroup.WithContext(ctx)
	return &Group{inner: inner, ctx: gctx}, gctx
}

func (g *Group) Wait() error {
	return g.inner.Wait()
}

func (g *Group) Go(f func() error) {
	g.inner.Go(func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				debug.PrintStack()
				if rerr, ok := r.(error); ok {
					err = fmt.Errorf("errgroup panic: %w", rerr)
				} else {
					err = fmt.Errorf("errgroup panic: %v", r)
				}
			}
		}()

		return f()
	})
}

func (g *Group) GoCtx(fn func(ctx context.Context) error) {
	g.Go(func() error {
		return fn(g.ctx)
	})
}
