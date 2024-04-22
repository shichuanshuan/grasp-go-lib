# 服务型 pprof 的应用
> 网页上要加上 url /debug/pprof

### 所有过去内存分配的采样
go tool pprof http://127.0.0.1:8080/debug/pprof/allocs

### 对活动对象的内存分配进行采样（活动）
go tool pprof http://127.0.0.1:8080/debug/pprof/heap

### 下载 cpu profile，默认从当前开始收集 30s 的 cpu 使用情况，需要等待 30s
go tool pprof http://127.0.0.1:8080/debug/pprof/profile

### wait 120s
go tool pprof http://127.0.0.1:8080/debug/pprof/profile?seconds=120

### 导致同步原语阻塞的堆栈跟踪
go tool pprof http://127.0.0.1:8080/debug/pprof/block

### 所有当前goroutine的堆栈跟踪
go tool pprof http://127.0.0.1:8080/debug/pprof/goroutine

### 争用互斥锁持有者的堆栈跟踪
go tool pprof http://127.0.0.1:8080/debug/pprof/mutex

### 当前程序的执行轨迹。
go tool pprof http://127.0.0.1:8080/debug/pprof/trace

原文链接：https://blog.csdn.net/weixin_38155824/article/details/124795704

# 下载的 pprof 通过 Web 访问
下载
运行程序后
curl http://127.0.0.1:8080/debug/pprof/goroutine?debug=1 -o goruntine.pprof

通过 Web 访问
go tool pprof -http=访问 IP go二进制文件 下载的pprof文件
go tool pprof -http=127.0.0.1:6600 /usr/local/go/bin/go profile