package main

import (
	"context"
	"fmt"
)

type key int

const nameKey key = 0

func worker(ctx context.Context) {
	if name, ok := ctx.Value(nameKey).(string); ok {
		fmt.Printf("worker: hello, %s!\n", name)
	} else {
		fmt.Println("worker: no name found")
	}
}

func main() {
	ctx := context.WithValue(context.Background(), nameKey, "Alice")

	go worker(ctx)

	// 等待一段时间，以便让 worker 完成工作
	fmt.Scanln()
}
