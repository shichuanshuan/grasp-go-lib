package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Second * 2)

	for tc := range tick {
		fmt.Println(tc)
	}
}
