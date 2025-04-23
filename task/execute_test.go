package task

import (
	"context"
	"errors"
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestSeqExecute(t *testing.T) {
	ctx := context.Background()

	{
		assert := utest.NewAssert(t, "one of tasks failed")
		tasks := []Task{
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return errors.New("err happeded") },
			func(ctx context.Context) error { return nil },
		}
		err := SeqExecute(ctx, tasks...)
		assert.IsNotNil(err)
	}

	{
		assert := utest.NewAssert(t, "all tasks executed successfully")
		tasks := []Task{
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return nil },
		}
		err := SeqExecute(ctx, tasks...)
		assert.IsNil(err)
	}
}

func TestConExecute(t *testing.T) {
	ctx := context.Background()

	{
		assert := utest.NewAssert(t, "one of tasks failed")
		tasks := []Task{
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return errors.New("err happeded") },
			func(ctx context.Context) error { return nil },
		}
		err := ConExecute(ctx, 10, 0, tasks...)
		assert.IsNotNil(err)
	}

	{
		assert := utest.NewAssert(t, "all tasks executed successfully")
		tasks := []Task{
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return nil },
		}
		err := ConExecute(ctx, 10, 0, tasks...)
		assert.IsNil(err)
	}
}

func TestSeqExecuteAll(t *testing.T) {
	ctx := context.Background()

	{
		assert := utest.NewAssert(t, "one of tasks failed")
		tasks := []Task{
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return errors.New("err happeded") },
			func(ctx context.Context) error { return nil },
		}
		err := SeqExecuteAll(ctx, tasks...)
		assert.Equal(len(err), 1)
	}
	{
		assert := utest.NewAssert(t, "all tasks failed")
		tasks := []Task{
			func(ctx context.Context) error { return errors.New("err happeded") },
			func(ctx context.Context) error { return errors.New("err happeded") },
			func(ctx context.Context) error { return errors.New("err happeded") },
		}
		err := SeqExecuteAll(ctx, tasks...)
		assert.Equal(len(err), len(tasks))
	}
	{
		assert := utest.NewAssert(t, "all tasks executed successfully")
		tasks := []Task{
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return nil },
		}
		errs := SeqExecuteAll(ctx, tasks...)
		assert.Equal(len(errs), 0)
	}
}

func TestConExecuteAll(t *testing.T) {
	ctx := context.Background()
	{
		assert := utest.NewAssert(t, "one of tasks failed")
		tasks := []Task{
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return errors.New("err happeded") },
			func(ctx context.Context) error { return nil },
		}
		err := ConExecuteAll(ctx, 10, tasks...)
		assert.Greater(len(err), 0)
	}

	{
		assert := utest.NewAssert(t, "all tasks failed")
		tasks := []Task{
			func(ctx context.Context) error { return errors.New("err happeded") },
			func(ctx context.Context) error { return errors.New("err happeded") },
			func(ctx context.Context) error { return errors.New("err happeded") },
		}
		err := ConExecuteAll(ctx, 10, tasks...)
		assert.Equal(len(err), len(tasks))
	}
}
