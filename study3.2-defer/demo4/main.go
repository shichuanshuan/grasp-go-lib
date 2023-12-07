package main

import (
	"fmt"
	"os"
	"time"
)

/*
上面的panic最终会被recover捕获到。这样的处理方式在一个http server的主流程常常会被用到。
一次偶然的请求可能会触发某个bug, 这时用recover捕获panic, 稳住主流程，不影响其他请求。
*/
func main() {
	defer fmt.Println("defer main")
	var user = os.Getenv("USER_")
	fmt.Println("user = ", user)

	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover success. err: ", err)
			}
		}()

		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == "" {
				panic("should set user env.")
			}

			// 此处不会执行
			fmt.Println("after panic")
		}()
	}()

	time.Sleep(100)
	fmt.Println("end of main function")
}
