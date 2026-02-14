package queue

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrQueueEmpty   = errors.New("queue is empty")
	ErrTaskNotFound = errors.New("task not found")
)

// MemoryQueue
type MemoryQueue struct {
	tasks []*Task
	mu    sync.Mutex
}

// NewMemoryQueue
func NewMemoryQueue() *MemoryQueue {
	return &MemoryQueue{
		tasks: make([]*Task, 0),
	}
}

// Enqueue
func (q *MemoryQueue) Enqueue(ctx context.Context, task *Task) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.tasks = append(q.tasks, task)
	return nil
}

// Dequeue
func (q *MemoryQueue) Dequeue(ctx context.Context) (*Task, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.tasks) == 0 {
		return nil, ErrQueueEmpty
	}

	task := q.tasks[0]
	q.tasks = q.tasks[1:]

	task.State = Processing
	return task, nil
}

// Ack
func (q *MemoryQueue) Ack(ctx context.Context, taskID string) error {
	return nil
}
