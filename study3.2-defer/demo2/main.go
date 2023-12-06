package main

import "fmt"

/*
理解这些坑的关键是这条语句：
return xxx
上面这条语句经过编译之后，变成了三条指令：

1. 返回值 = xxx
2. 调用defer函数
3. 空的return
*/

// 拆解之前
func f() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

// 拆解之后
/*
func f() (r int) {
	t := 5

	// 1. 赋值指令
	r = t

	// 2. defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
	func() {
		t = t + 5
	}

	// 3. 空的return指令
	return
}
*/

func main() {
	num := f()
	fmt.Println(num)
}
