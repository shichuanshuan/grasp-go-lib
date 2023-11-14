package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Printf("worker: deadline set to %s\n", deadline.Format(time.RFC3339))
	}

	select {
	case <-time.After(time.Second * 2):
		fmt.Println("worker completed")
	case <-ctx.Done():
		fmt.Println("worker cancelled")
	}
}

func main() {
	d := time.Now().Add(time.Second * 3)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	go worker(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("main cancelled")
		time.Sleep(1 * time.Second)
	case <-time.After(time.Second * 4):
		fmt.Println("main completed")
	}
}
