package main

import (
	"fmt"
	"time"
)

func main() {
	now := <-time.After(time.Second * 2)

	fmt.Println(now)
}
