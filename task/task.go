package task

import (
	"context"
	"fmt"
)

// AsyncResult 异步执行函数的结果，包含结果值和错误
type AsyncResult[T any] struct {
	Value T
	Err   error
}

// AsyncExecute executes function asynchronously.
func AsyncExecute[T any](ctx context.Context, fn func(ctx context.Context) (T, error)) <-chan AsyncResult[T] {
	resultCh := make(chan AsyncResult[T], 1)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				resultCh <- AsyncResult[T]{Err: fmt.Errorf("panic: %v", e)}
			}
			close(resultCh)
		}()
		select {
		// Check if ctx has timed out or been cancelled.
		case <-ctx.Done():
			resultCh <- AsyncResult[T]{Err: ctx.Err()}
		// Execute asynchronous function.
		default:
			value, err := fn(ctx)
			resultCh <- AsyncResult[T]{Value: value, Err: err}
		}
	}()
	return resultCh
}
