package main

import (
	"context"
	"log"
	"net"

	pb "github.com/kmranimesh/titan/api/proto"
	"github.com/kmranimesh/titan/internal/queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server
type server struct {
	pb.UnimplementedTitanServer
	q queue.Queue
}

// Enqueue
func (s *server) Enqueue(ctx context.Context, req *pb.EnqueueRequest) (*pb.EnqueueResponse, error) {
	task := queue.NewTask(req.Type, req.Payload)
	if err := s.q.Enqueue(ctx, task); err != nil {
		return nil, err
	}
	log.Printf("Enqueued task: %s", task.ID)
	return &pb.EnqueueResponse{Id: task.ID}, nil
}

// Poll
func (s *server) Poll(ctx context.Context, req *pb.PollRequest) (*pb.PollResponse, error) {
	task, err := s.q.Dequeue(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("Polled task: %s", task.ID)
	return &pb.PollResponse{
		Id:      task.ID,
		Type:    task.Type,
		Payload: task.Payload,
	}, nil
}

// Ack
func (s *server) Ack(ctx context.Context, req *pb.AckRequest) (*pb.AckResponse, error) {
	if err := s.q.Ack(ctx, req.Id); err != nil {
		return nil, err
	}
	log.Printf("Acked task: %s", req.Id)
	return &pb.AckResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	
	// Register our implementation
	// Note: We use &server{} because our methods are defined on the pointer receiver
	pb.RegisterTitanServer(s, &server{
		q: queue.NewMemoryQueue(),
	})
	
	// Enable reflection for CLI tools (like evilplot/grpcurl)
	reflection.Register(s)

	log.Printf("Titan Broker listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
