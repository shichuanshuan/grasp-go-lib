package main

import "internal/testlog"

func main() {
	// 搞懂为什么不能调用
	testlog.Getenv("python")
}
