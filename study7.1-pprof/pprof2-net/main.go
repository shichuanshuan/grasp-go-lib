package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	//用于pprof检测内存使用情况
	go http.ListenAndServe("0.0.0.0:8080", nil)

	sum := 0
	for {

		go func() {
			fmt.Println("sum= ", sum)
			time.Sleep(time.Minute)
		}()
		sum += 1
		time.Sleep(time.Second)
		fmt.Println("sum= ", sum)
	}

}

/*
#所有过去内存分配的采样
go tool pprof http://127.0.0.1:8080/debug/pprof/allocs

#对活动对象的内存分配进行采样（活动）
go tool pprof http://127.0.0.1:8080/debug/pprof/heap

# 下载 cpu profile，默认从当前开始收集 30s 的 cpu 使用情况，需要等待 30s
go tool pprof http://127.0.0.1:8080/debug/pprof/profile
# wait 120s
go tool pprof http://127.0.0.1:8080/debug/pprof/profile?seconds=120

#导致同步原语阻塞的堆栈跟踪
go tool pprof http://127.0.0.1:8080/debug/pprof/block

#所有当前goroutine的堆栈跟踪
go tool pprof http://127.0.0.1:8080/debug/pprof/goroutine

#争用互斥锁持有者的堆栈跟踪
go tool pprof http://127.0.0.1:8080/debug/pprof/mutex

#当前程序的执行轨迹。
go tool pprof http://127.0.0.1:8080/debug/pprof/trace
*/
