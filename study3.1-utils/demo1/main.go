package main

import (
	"demo1/utils"
	"fmt"
)

func hello() {
	buf := utils.GetBuffer()
	defer utils.PutBuffer(buf)

	buf.Data = []byte("hello world")
}

func readData() {
	buf := utils.GetBuffer()
	defer utils.PutBuffer(buf)

	buf.Data = []byte("shics")

}

func main() {

	// 存一个数据
	hello()
	// 读一个数据
	buf := utils.GetBuffer()
	fmt.Printf("buf [%s]\n", buf.Data)

	// 存一个数据
	readData()
	// 读一个数据
	buf = utils.GetBuffer()
	fmt.Printf("read %s\n", buf.Data)
}
