package main

import (
	"demo2/utils"
	"fmt"
)

func hello() {
	buf := utils.GetBuffer()

	buf.Data = []byte("hello world")
	utils.PutBuffer(buf)
}

func readData() {
	//buf := utils.GetBuffer()
	buf := utils.Buffer{}

	buf.Data = []byte("shics")
	utils.PutBuffer(&buf)
}

func main() {

	// 存一个数据
	hello()
	readData()

	// 读一个数据
	buf := utils.GetBuffer()
	fmt.Printf("buf [%s]\n", buf.Data)

	// 读一个数据
	buf = utils.GetBuffer()
	fmt.Printf("read [%s]\n", buf.Data)
}
