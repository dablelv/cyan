package task

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

// AsyncResult is the result of the function executed asynchronously.
type AsyncResult[T any] struct {
	Val T
	Err error
}

// AsyncExecute executes function asynchronously.
func AsyncExecute[T any](ctx context.Context, fn func(ctx context.Context) (T, error)) <-chan AsyncResult[T] {
	rsltCh := make(chan AsyncResult[T], 1)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				rsltCh <- AsyncResult[T]{Err: fmt.Errorf("panic: %v", e)}
			}
			close(rsltCh)
		}()
		select {
		// Check if ctx has timed out or been cancelled.
		case <-ctx.Done():
			rsltCh <- AsyncResult[T]{Err: ctx.Err()}
		// Execute asynchronous function.
		default:
			v, err := fn(ctx)
			rsltCh <- AsyncResult[T]{Val: v, Err: err}
		}
	}()
	return rsltCh
}

// BatchConcurrent batch execute function concurrently.
// The concurrency indicates how many fn are executing at the same time.
func BatchConcurrent[T, R any](ctx context.Context, concurrency int, slice []T, fn func(context.Context, T) (R, error)) ([]R, error) {
	var eg errgroup.Group
	r := make([]R, len(slice))

	for i := 0; i < len(slice); {
		for j := 0; j < concurrency; j++ {
			idx := i + j

			// The last batch less than the specified quantity.
			if idx == len(slice) {
				break
			}
			eg.Go(func() (err error) {
				defer func() {
					if e := recover(); e != nil {
						err = fmt.Errorf("panic: %w", e.(error))
					}
				}()
				r[idx], err = fn(ctx, slice[idx])
				return
			})
		}
		if err := eg.Wait(); err != nil {
			return nil, err
		}
		// Next batch.
		i += concurrency
	}
	return r, nil
}
