package task

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/dablelv/cyan/internal"
)

func TestDoWithRetry(t *testing.T) {
	ctx := context.Background()

	assert1 := internal.NewAssert(t, "f return success, should no retry")
	f1 := func() error { return nil }
	err := DoWithRetry(ctx, f1, 3, 10*time.Millisecond)
	assert1.IsNil(err)

	assert2 := internal.NewAssert(t, "f always fail, should return error")
	f2 := func() error { return errors.New("some error") }
	err = DoWithRetry(ctx, f2, 3, 10*time.Millisecond)
	assert2.ErrorContains(err, "some error")

	assert3 := internal.NewAssert(t, "context cancelled, should return context canceled")
	cancelctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-time.After(10 * time.Millisecond)
		cancel()
	}()
	f3 := func() error { return errors.New("some error") }
	err = DoWithRetry(cancelctx, f3, 3, 10*time.Millisecond)
	assert3.ErrorContains(err, "context canceled")
}
