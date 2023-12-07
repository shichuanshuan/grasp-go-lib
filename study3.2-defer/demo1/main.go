package main

import (
	"fmt"
)

type number int

func (n number) print()   { fmt.Println(n) }
func (n *number) pprint() { fmt.Println(*n) }

/*
在defer函数定义时，对外部变量的引用是有两种方式的，分别是作为函数参数和作为闭包引用。
作为函数参数，则在defer定义时就把值传递给defer，并被cache起来；
作为闭包引用的话，则会在defer函数真正调用时根据整个上下文确定当前的值。
*/
func main() {
	var n number

	defer n.print()
	defer n.pprint()
	defer func() { n.print() }()
	defer func() { n.pprint() }()

	n = 3

}

/*
第四个defer语句是闭包，引用外部函数的n, 最终结果是3;
第三个defer语句同第四个；
第二个defer语句，n是引用，最终求值是3.
第一个defer语句，对n直接求值，开始的时候n=0, 所以最后是0;
*/
