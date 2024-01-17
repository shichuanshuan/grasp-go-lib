package main

import (
	"fmt"
	"sync"
	"time"
)
// RWMutex是对Mutex的扩展，允许多个goroutine同时读取共享资源，但在写入时会互斥。这在读多写少的场景中特别有用。下面是一个示例：
var (
	data    map[string]string
	rwMutex sync.RWMutex
)

func readData(key string) string {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return data[key]
}

func writeData(key, value string) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	data[key] = value
}

func main() {
	data = make(map[string]string)

	var wg sync.WaitGroup

	// 写入数据
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", index)
			value := fmt.Sprintf("value%d", index)
			writeData(key, value)
			time.Sleep(time.Millisecond * 100) // 模拟写入耗时
		}(i)
	}

	// 读取数据
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", index%5)
			value := readData(key)
			fmt.Printf("Read: %s - %s\n", key, value)
		}(i)
	}

	wg.Wait()
}
// 在上述例子中，RWMutex用于保护对data的读写操作。多个goroutine可以同时读取，但在写入时会相互互斥。
