package main

import (
	"fmt"
	"sync"
	"time"
)
// WaitGroup用于等待一组goroutine完成执行。它在主goroutine中等待所有其他goroutine执行完成后再继续执行。下面是一个简单的例子：
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers done")
}
// 在上面的例子中，WaitGroup确保所有的worker goroutine都执行完成后，才会继续执行主goroutine。