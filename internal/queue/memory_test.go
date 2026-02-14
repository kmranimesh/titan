package queue

import (
	"context"
	"testing"
)

// TestMemoryQueue_EnqueueDequeue
func TestMemoryQueue_EnqueueDequeue(t *testing.T) {
	q := NewMemoryQueue()
	ctx := context.Background()

	task := NewTask("email:send", []byte("test"))

	if err := q.Enqueue(ctx, task); err != nil {
		t.Fatalf("Enqueue failed: %v", err)
	}

	dequeued, err := q.Dequeue(ctx)
	if err != nil {
		t.Fatalf("Dequeue failed: %v", err)
	}

	if dequeued.ID != task.ID {
		t.Errorf("Expected task ID %s, got %s", task.ID, dequeued.ID)
	}
}

// TestMemoryQueue_Empty
func TestMemoryQueue_Empty(t *testing.T) {
	q := NewMemoryQueue()
	ctx := context.Background()

	_, err := q.Dequeue(ctx)
	if err != ErrQueueEmpty {
		t.Errorf("Expected ErrQueueEmpty, got %v", err)
	}
}
