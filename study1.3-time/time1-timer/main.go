package main

import (
	"fmt"
	"sync"
	"time"
)

// 多线程示例
func manyGo() {
	var wg sync.WaitGroup
	timer := time.NewTimer(2 * time.Second) // 创建一个定时器，2秒后触发
	defer timer.Stop()                      // 停止定时器的执行

	fmt.Println("Start")

	// 启动3个协程
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 等待定时器到期
			<-timer.C
			fmt.Printf("Goroutine %d: Timer expired\n", id)
		}(i)
	}

	// 让主协程休眠一段时间，以便其他协程有机会等待定时器
	time.Sleep(3 * time.Second)

	// 停止定时器
	if !timer.Stop() {
		<-timer.C // 如果定时器已经触发，从通道中读取
	}

	wg.Wait() // 等待所有协程完成
	fmt.Println("All goroutines finished")
}

func main1() {
	fmt.Println("start")
	d := time.Second * 2

	// 创建一个计时器，指定触发时间为 2 秒
	tm := time.NewTimer(d)

	// 阻塞等待计时器到期
	<-tm.C

	// 计时器不使用时，关闭计时器，避免过期
	ret := tm.Stop()
	fmt.Println("stop ret:", ret)
}
