package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	d := time.Second * 2

	// 创建一个计时器，指定触发时间未 2 秒
	tm := time.NewTimer(d)

	// 阻塞等待计时器到期
	<-tm.C

	// 计时器不使用时，关闭计时器，避免过期
	ret := tm.Stop()
	fmt.Println("stop ret:", ret)
}
