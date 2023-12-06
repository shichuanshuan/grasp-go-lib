package main

import (
	"fmt"
	"io"
	"os"
)

/*
在一个函数里，需要打开两个文件进行合并操作，合并完后，在函数执行完后关闭打开的文件句柄。
*/
func mergeFile() error {
	// 使用的是同一个变量名 f，defer 复制的是当前的值
	f, _ := os.Open("file1.txt")
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close file1.txt err %v\n", err)
			}
		}(f)
	}

	// ……

	f, _ = os.Open("file2.txt")
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close file2.txt err %v\n", err)
			}
		}(f)
	}

	return nil
}

/*
上面的代码中就用到了defer的原理，defer函数定义的时候，参数就已经复制进去了，之后，真正执行close()函数的时候就刚好关闭的是正确的“文件”了，妙哉！
可以想像一下如果不这样将f当成函数参数传递进去的话，最后两个语句关闭的就是同一个文件了，都是最后一个打开的文件。

不过在调用close()函数的时候，要注意一点：先判断调用主体是否为空，否则会panic. 比如上面的代码片段里，先判断f不为空，才会调用Close()函数，这样最安全。
*/
func main() {
	err := mergeFile()
	fmt.Println(err)
}
