package main

import (
	"fmt"
	"sync"
	"time"
)
// Mutex是最基本的同步原语之一，用于保护共享资源，确保在同一时刻只有一个goroutine可以访问它。下面是一个简单的示例：
var counter int
var mutex sync.Mutex

func increment() {
	mutex.Lock()
	defer mutex.Unlock()
	counter++
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
//在上面的例子中，Mutex用于锁定counter变量，以确保在任何时刻只有一个goroutine能够递增它。WaitGroup用于等待所有goroutine完成后再继续执行主程序。