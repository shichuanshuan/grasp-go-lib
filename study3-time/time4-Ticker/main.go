package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.NewTicker(time.Second * 1)

	go func() {
		time.Sleep(time.Second * 5)
		tick.Stop()
	}()

	for tc := range tick.C {
		fmt.Println("tc =", tc)
	}

	fmt.Println("--------")
	// Tick 效果等同于 time.NewTicker()
	ntc := time.Tick(time.Second * 2)
	for tck := range ntc {
		fmt.Println(tck)
	}

}
