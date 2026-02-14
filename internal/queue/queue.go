package queue

import "context"

// Queue
type Queue interface {
	Enqueue(ctx context.Context, task *Task) error
	Dequeue(ctx context.Context) (*Task, error)
	Ack(ctx context.Context, taskID string) error
}
