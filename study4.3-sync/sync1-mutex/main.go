package main

import (
	"fmt"
	"sync"
	"time"
)

// Mutex是最基本的同步原语之一，用于保护共享资源，确保在同一时刻只有一个goroutine可以访问它。下面是一个简单的示例：
var counter int
var mutex sync.Mutex

// 因为
func increment() {
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup

	start := time.Now().UnixNano()

	//起了10000个协程让sum自增1，但是结果却每次不一样，因为竞态，多个协程可能获取了同一个值，因此出现问题
	//使用sync.Mutex互斥锁解决此问题
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	end := time.Now().UnixNano()
	fmt.Println("Counter:", counter)
	fmt.Printf("timer %v\n", end-start)
}
