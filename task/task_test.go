package task

import (
	"context"
	"errors"
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestAsyncExecute(t *testing.T) {
	ctx := context.Background()

	{
		assert := utest.NewAssert(t, "success")
		fn := func(ctx context.Context) (any, error) {
			return "success", nil
		}
		rc := AsyncExecute(ctx, fn)
		r := <-rc
		assert.IsNil(r.Err)
		assert.Equal("success", r.Val)
	}

	{
		assert := utest.NewAssert(t, "error")
		fn := func(ctx context.Context) (any, error) {
			return "", errors.New("an error")
		}
		rc := AsyncExecute(ctx, fn)
		r := <-rc
		assert.IsNotNil(r.Err)
	}

	{
		assert := utest.NewAssert(t, "panic")
		fn := func(ctx context.Context) (any, error) {
			panic("a panic")
		}
		rc := AsyncExecute(ctx, fn)
		r := <-rc
		assert.IsNotNil(r.Err)
	}
}

func TestBatchConcurrent(t *testing.T) {
	ctx := context.Background()

	{
		assert := utest.NewAssert(t, "success")
		ns := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		concurrency := 4
		fn := func(ctx context.Context, i int) (int, error) {
			return i * 2, nil
		}
		r, err := BatchConcurrent(ctx, concurrency, ns, fn)
		assert.IsNil(err)
		assert.Equal(r, []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20})
	}

	{
		assert := utest.NewAssert(t, "error")
		ns := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		concurrency := 4
		fn := func(ctx context.Context, i int) (int, error) {
			if i == 4 {
				return 0, errors.New("an error")
			}
			return i * 2, nil
		}
		r, err := BatchConcurrent(ctx, concurrency, ns, fn)
		assert.ErrorContains(err, "error")
		assert.IsNil(r)
	}

	{
		assert := utest.NewAssert(t, "panic")
		ns := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		concurrency := 4
		fn := func(ctx context.Context, i int) (int, error) {
			if i == 4 {
				panic(errors.New("a panic"))
			}
			return i * 2, nil
		}
		r, err := BatchConcurrent(ctx, concurrency, ns, fn)
		assert.ErrorContains(err, "panic")
		assert.IsNil(r)
	}
}
