package main

import (
	"errors"
	"fmt"
)

func f1() {
	var err error

	defer fmt.Println(err)

	err = errors.New("defer error")
	return
}

func f2() {
	var err error

	defer func() {
		fmt.Println(err)
	}()

	err = errors.New("defer error")
	return
}

func f3() {
	var err error

	defer func(err error) {
		fmt.Println(err)
	}(err)

	err = errors.New("defer error")
	return
}

/*
第1，3个函数是因为作为函数参数，定义的时候就会求值，定义的时候err变量的值都是nil, 所以最后打印的时候都是nil.
第2个函数的参数其实也是会在定义的时候求值，只不过，第2个例子中是一个闭包，它引用的变量err在执行的时候最终变成defer error了。
关于闭包在本文后面有介绍。
*/
func main() {
	f1()
	f2()
	f3()
}
