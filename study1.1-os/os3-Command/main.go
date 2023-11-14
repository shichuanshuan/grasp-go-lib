package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
)

func main() {
	ps := exec.Command("ps", "-ef")
	grep := exec.Command("grep", "-i", "ssh")

	r, w := io.Pipe() // 创建一个管道
	defer r.Close()
	defer w.Close()
	ps.Stdout = w  // ps向管道的一端写
	grep.Stdin = r // grep从管道的一端读

	var buffer bytes.Buffer
	grep.Stdout = &buffer // grep的输出为buffer

	_ = ps.Start()
	_ = grep.Start()
	ps.Wait()
	w.Close()
	grep.Wait()
	io.Copy(os.Stdout, &buffer) // buffer拷贝到系统标准输出
}
