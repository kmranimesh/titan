package worker

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/kmranimesh/titan/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Handler
type Handler func(ctx context.Context, taskType string, payload []byte) error

// Worker
type Worker struct {
	client     pb.TitanServerClient
	conn       *grpc.ClientConn
	handlers   map[string]Handler
	brokerAddr string
}

// NewWorker
func NewWorker(brokerAddr string) (*Worker, error) {
	conn, err := grpc.Dial(brokerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to broker: %w", err)
	}

	return &Worker{
		client:     pb.NewTitanServerClient(conn),
		conn:       conn,
		handlers:   make(map[string]Handler),
		brokerAddr: brokerAddr,
	}, nil
}

// Register
func (w *Worker) Register(taskType string, handler Handler) {
	w.handlers[taskType] = handler
}

// Start
func (w *Worker) Start(ctx context.Context) error {
	log.Printf("Worker started, connected to %s", w.brokerAddr)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			resp, err := w.client.Poll(ctx, &pb.PollRequest{})
			if err != nil {
				continue
			}

			if handler, ok := w.handlers[resp.Type]; ok {
				if err := handler(ctx, resp.Type, resp.Payload); err == nil {
					w.client.Ack(ctx, &pb.AckRequest{Id: resp.Id})
				} else {
					log.Printf("Task %s failed: %v", resp.Id, err)
				}
			} else {
				log.Printf("No handler for task type: %s", resp.Type)
			}
		}
	}
}

// Close
func (w *Worker) Close() error {
	return w.conn.Close()
}
