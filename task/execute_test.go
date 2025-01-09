package task

import (
	"context"
	"errors"
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestExecuteTasks(t *testing.T) {
	ctx := context.Background()

	{
		assert := utest.NewAssert(t, "any task failed")
		tasks := []Task{
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return errors.New("err happeded") },
			func(ctx context.Context) error { return nil },
		}
		err := ExecuteTasks(ctx, tasks)
		assert.IsNotNil(err)
	}

	{
		assert := utest.NewAssert(t, "all tasks executed successfully")
		tasks := []Task{
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return nil },
		}
		err := ExecuteTasks(ctx, tasks)
		assert.IsNil(err)
	}
}
