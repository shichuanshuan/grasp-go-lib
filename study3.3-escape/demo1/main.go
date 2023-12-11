package main

import "fmt"

func foo() *int {
	t := 3
	return &t
}

/*
foo函数返回一个局部变量的指针，main函数里变量x接收它。执行如下命令：
go build -gcflags '-m -l' main.go
加-l是为了不让foo函数被内联。
*/
func main() {
	x := foo()
	fmt.Println(*x)
}

/*
让我们不解的是为什么main函数里的x也逃逸了？这是因为有些函数参数为interface类型，
比如fmt.Println(a ...interface{})，编译期间很难确定其参数的具体类型，也会发生逃逸。
*/
