package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kmranimesh/titan/pkg/client"
	"github.com/kmranimesh/titan/pkg/worker"
)

func main() {
	// Worker
	w, err := worker.NewWorker("localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	w.Register("email:send", func(ctx context.Context, taskType string, payload []byte) error {
		fmt.Printf(" [Worker] Processing task: %s | Payload: %s\n", taskType, string(payload))
		return nil
	})

	go func() {
		if err := w.Start(context.Background()); err != nil {
			log.Printf("Worker error: %v", err)
		}
	}()

	// Producer
	p, err := client.NewProducer("localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		id, err := p.Enqueue(context.Background(), "email:send", []byte(fmt.Sprintf("Hello Titan %d", i)))
		if err != nil {
			log.Printf("Enqueue failed: %v", err)
		} else {
			fmt.Printf(" [Producer] Enqueued task: %s\n", id)
		}
		time.Sleep(500 * time.Millisecond)
	}

	time.Sleep(2 * time.Second)
}
