package task

import "context"

// Task is a function that takes a context and returns an error.
type Task func(ctx context.Context) error

// ExecuteTasks execute tasks sequentially.
func ExecuteTasks(ctx context.Context, tasks []Task) error {
	for _, task := range tasks {
		if err := task(ctx); err != nil {
			return err
		}
	}
	return nil
}
