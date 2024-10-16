package task

import (
	"context"
	"math"
	"time"
)

func DoWithRetry(ctx context.Context, f func() error, maxAttempts int, backoff time.Duration) error {
	var err error
	for attempts := 0; attempts < maxAttempts; attempts++ {
		if err = f(); err == nil {
			return nil
		}
		// No more retries after the maximum number of retries.
		if attempts == maxAttempts-1 {
			break
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Duration(math.Pow(2, float64(attempts))) * backoff):
		}
	}
	// If the operation still fails after the maximum number of retries, the last error is returned.
	return err
}
