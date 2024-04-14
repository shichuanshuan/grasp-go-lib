package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// 启动 HTTP 服务器，监听 6060 端口
	go func() {
		log.Println(http.ListenAndServe("127.0.0.1:6060", nil))
	}()

	// 你的其他服务器逻辑代码
	// 这里只是一个示例，你可以将你的服务器代码放在这里
	log.Println("Server started.")

	// 这里让程序保持运行，以保持 HTTP 服务器服务
	select {}
}
