package main

import (
	"fmt"
	"sync"
	"time"
)

// Cond提供了一种在goroutine之间等待或发信号的机制。下面是一个使用Cond的例子：
var (
	count    int
	maxCount = 3
	cond     *sync.Cond
)

func producer() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		cond.L.Lock()
		count++
		fmt.Printf("Produced: %d\n", count)
		cond.Signal() // 发送信号通知消费者
		cond.L.Unlock()
	}
}

func consumer() {
	for {
		cond.L.Lock()
		for count < 1 {
			fmt.Println("wait...", count, count < 1)
			cond.Wait() // 等待生产者发信号
		}
		fmt.Printf("Consumed: %d\n", count)
		count--
		cond.L.Unlock()
	}
}

// Cond用于在生产者和消费者之间同步数据。生产者每秒生产一个数据，而消费者则等待信号并消费数据。
func main() {
	cond = sync.NewCond(&sync.Mutex{})
	go producer()
	go consumer()

	time.Sleep(10 * time.Second)
}
