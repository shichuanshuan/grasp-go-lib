package main

import (
	"fmt"
	"internal/godebug"
	"net/http"
)

var enableDebug = godebug.New("enableDebug")

func main() {
	// 从环境变量中获取 enableDebug 的值
	fmt.Println("enableDebug:", enableDebug.Value())

	// 启动 HTTP 服务器
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if enableDebug.Value() == "1" {
		// 如果 enableDebug 为 "1"，输出调试信息
		fmt.Println("Handling request:", r.URL.Path)
	} else {
		fmt.Fprintf(w, "Hello, World!")
	}
}
