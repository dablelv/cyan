package task

import "context"

// Task is a function that takes a context and returns an error.
type Task func(ctx context.Context) error

// ExecTasks execute tasks sequentially.
func ExecTasks(ctx context.Context, tasks []Task) error {
	for _, task := range tasks {
		if err := task(ctx); err != nil {
			return err
		}
	}
	return nil
}
