package main

import (
	"fmt"
	"sync"
	"time"
)
// Mutex���������ͬ��ԭ��֮һ�����ڱ���������Դ��ȷ����ͬһʱ��ֻ��һ��goroutine���Է�������������һ���򵥵�ʾ����
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
//������������У�Mutex��������counter��������ȷ�����κ�ʱ��ֻ��һ��goroutine�ܹ���������WaitGroup���ڵȴ�����goroutine��ɺ��ټ���ִ��������