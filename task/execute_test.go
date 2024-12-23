package task

import (
	"context"
	"errors"
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestExecuteTasks(t *testing.T) {
	ctx := context.Background()

	func() {
		assert := utest.NewAssert(t, "any task failed")
		tasks := []Task{
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return errors.New("err happeded") },
			func(ctx context.Context) error { return nil },
		}
		err := ExecTasks(ctx, tasks)
		assert.IsNotNil(err)
	}()

	func() {
		tasks := []Task{
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return nil },
		}
		err := ExecTasks(ctx, tasks)
		assert := utest.NewAssert(t, "all tasks executed successfully")
		assert.IsNil(err)
	}()
}
