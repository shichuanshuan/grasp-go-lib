package main

import "fmt"

type study struct {
	name string
	age  int
}

func main() {
	s := study{name: "shi", age: 25}
	fmt.Println("shics")
	fmt.Println("name", s)
}
