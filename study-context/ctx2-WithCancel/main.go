package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {

	select {
	default:
		fmt.Printf("worker %d is running\n", id)
	case <-ctx.Done():
		fmt.Printf("worker %d is cancelled\n", id)
		return
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动两个 worker
	go worker(ctx, 1)
	go worker(ctx, 2)

	// 运行一段时间后取消所有 worker
	time.Sleep(time.Second * 3)
}
