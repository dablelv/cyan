package safegroup

import (
	"context"
	"errors"
	"fmt"
	"runtime"

	"golang.org/x/sync/errgroup"
)

// PanicBufLen panic stack buffer max size.
var PanicBufLen = 2048

// Group secure group.
type Group struct {
	g errgroup.Group
}

// WithContext returns a new Group and an associated Context derived from ctx.
func WithContext(ctx context.Context) (*Group, context.Context) {
	g, ctx := errgroup.WithContext(ctx)
	return &Group{g: *g}, ctx
}

// Go calls the given function in a new goroutine.
// If a panic occurred in a given function, it will be captured and error will be returned.
// At the same time, the goroutine stack information will be printed to standard output.
func (s *Group) Go(f func() error) {
	s.g.Go(func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, PanicBufLen)
				buf = buf[:runtime.Stack(buf, false)]
				fmt.Printf("[PANIC]:\n%+v\n[STACK]:\n%s\n", r, buf)
				err = errors.New(fmt.Sprint(r))
			}
		}()
		return f()
	})
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (s *Group) Wait() error {
	return s.g.Wait()
}
