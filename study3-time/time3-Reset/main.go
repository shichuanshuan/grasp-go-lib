package main

import (
	"fmt"
	"time"
)

func main() {
	tm := time.NewTimer(time.Second * 1)

	fmt.Println("start")
	res := tm.Reset(time.Second * 2)
	fmt.Println("res =", res)

	<-tm.C

	fmt.Println("start 2")
	res2 := tm.Reset(time.Second * 10)
	fmt.Println("ret2 ", res2)

	<-tm.C

	fmt.Println("end")
}
