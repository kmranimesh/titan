package client

import (
	"context"
	"fmt"

	pb "github.com/kmranimesh/titan/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Producer
type Producer struct {
	client pb.TitanServerClient
	conn   *grpc.ClientConn
}

// NewProducer
func NewProducer(brokerAddr string) (*Producer, error) {
	conn, err := grpc.Dial(brokerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to broker: %w", err)
	}

	return &Producer{
		client: pb.NewTitanServerClient(conn),
		conn:   conn,
	}, nil
}

// Close
func (p *Producer) Close() error {
	return p.conn.Close()
}

// Enqueue
func (p *Producer) Enqueue(ctx context.Context, taskType string, payload []byte) (string, error) {
	resp, err := p.client.Enqueue(ctx, &pb.EnqueueRequest{
		Type:    taskType,
		Payload: payload,
	})
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}
