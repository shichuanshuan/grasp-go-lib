package main

import (
	"demo2/utils"
	"fmt"
)

func hello() {
	buf := utils.GetBuffer()
	defer utils.PutBuffer(buf)

	buf.Data = []byte("hello world")
}

func readData() {
	//buf := utils.GetBuffer()
	buf := utils.Buffer{}
	defer utils.PutBuffer(&buf)

	buf.Data = []byte("shics")

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
