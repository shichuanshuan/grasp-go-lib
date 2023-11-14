package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个带有取消功能的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Second)
	defer cancel()
	// 模拟一个耗时操作，在 3 秒后取消上下文

	time.Sleep(3 * time.Second)

	// 执行操作并检查上下文状态
	if ctx != nil {
		select {
		case <-ctx.Done():
			fmt.Println("操作被取消:", ctx.Err())
		default:
			fmt.Println("操作正常完成")
			// 这里可以添加具体的操作代码
		}
	} else {
		fmt.Println("上下文为空")
	}
}
