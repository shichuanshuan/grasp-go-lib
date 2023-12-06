package main

import "fmt"

func f() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

/*
拆解后：

func f() (r int) {
     // 1. 赋值
     r = 1

     // 2. 这里改的r是之前传值传进去的r，不会改变要返回的那个r值
     func(r int) {
          r = r + 5
     }(r)

     // 3. 空的return
     return
}
*/

func main() {
	num := f()
	fmt.Println(num)
}
