package task

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// Task is a function that takes a context and returns an error.
type Task func(ctx context.Context) error

// SeqExecute executes tasks sequentially.
// If an error is encountered, return immediately.
func SeqExecute(ctx context.Context, tasks ...Task) error {
	for _, task := range tasks {
		if err := task(ctx); err != nil {
			return err
		}
	}
	return nil
}

// ConExecute concurrently execute tasks with concurrent upper limit.
// If timeout is 0, there will be no timeout.
// If any one of tasks reports an error and the others are interrupted immediately.
func ConExecute(ctx context.Context, limit int, timeout time.Duration, tasks ...Task) error {
	var timeCh <-chan time.Time
	if timeout > 0 {
		timeCh = time.After(timeout)
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(limit)

	// When any goroutine returns an error, cancel other running goroutines by listening to ctx.Done() immediately
	// and return the first non-nil error in the Wait method.
	for _, task := range tasks {
		fn := task
		eg.Go(func() (err error) {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-timeCh:
				return fmt.Errorf("execution timeout %v", timeout)
			default:
				defer func() {
					if panicErr := recover(); panicErr != nil {
						err = fmt.Errorf("execution panic: %w", panicErr.(error))
					}
				}()
				return fn(ctx)
			}
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// SeqExecuteAll executes all tasks sequentially.
// If one of tasks reports an error, continue to execute the remaining tasks and return all errors occurred.
func SeqExecuteAll(ctx context.Context, tasks ...Task) (errs []error) {
	for _, task := range tasks {
		if err := task(ctx); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// ConExecuteAll concurrently executes all tasks with concurrent upper limit.
// If one of tasks reports an error, continue to execute the remaining tasks and return all errors occurred.
func ConExecuteAll(ctx context.Context, limit int, tasks ...Task) (errs []error) {
	var eg errgroup.Group

	// Control the number of concurrent goroutines.
	eg.SetLimit(limit)

	errCh := make(chan error, len(tasks))
	for _, task := range tasks {
		fn := task
		eg.Go(func() error {
			defer func() {
				if panicErr := recover(); panicErr != nil {
					errCh <- fmt.Errorf("execution panic: %w", panicErr.(error))
				}
			}()
			if err := fn(ctx); err != nil {
				errCh <- err
			}
			return nil
		})
	}
	eg.Wait()
	close(errCh)
	for chErr := range errCh {
		errs = append(errs, chErr)
	}
	return
}
