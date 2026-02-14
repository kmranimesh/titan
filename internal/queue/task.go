package queue

import (
	"time"

	"github.com/google/uuid"
)

// State
type State int

const (
	Pending State = iota
	Processing
	Completed
	Failed
)

func (s State) String() string {
	return [...]string{"Pending", "Processing", "Completed", "Failed"}[s]
}

// Task
type Task struct {
	ID        string
	Type      string
	Payload   []byte
	State     State
	CreatedAt time.Time
}

// NewTask
func NewTask(taskType string, payload []byte) *Task {
	return &Task{
		ID:        uuid.NewString(),
		Type:      taskType,
		Payload:   payload,
		State:     Pending,
		CreatedAt: time.Now(),
	}
}
