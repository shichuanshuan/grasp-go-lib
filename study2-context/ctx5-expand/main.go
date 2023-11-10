package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 创建一个父Context和取消函数
	ctx, cancel := context.WithCancel(context.Background())

	// 创建一个通道用于接收操作系统的终止信号
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	// 启动多个子进程
	for i := 1; i <= 5; i++ {
		go runChildProcess(ctx, i)
	}

	// 等待操作系统的终止信号
	<-sigCh

	// 收到终止信号后，取消父Context，所有子进程也会被取消
	cancel()

	// 等待子进程完成
	time.Sleep(time.Second)

	fmt.Println("All processes have been stopped")
}

func runChildProcess(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			// 父Context被取消，表示要停止子进程
			fmt.Println("Process", id, "is stopping")
			return
		default:
			// 子进程的操作逻辑
			fmt.Println("Process", id, "is running")
			time.Sleep(time.Second)
		}
	}
}
