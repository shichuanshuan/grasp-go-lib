package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	// 2 秒后执行回调函数
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("Timer expired")
	})

	// 两秒过后，匿名函数还需要时间执行，因此要等待
	time.Sleep(3 * time.Second)
}
