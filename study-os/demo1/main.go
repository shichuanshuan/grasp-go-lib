package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 后缀以 .COM;.EXE;.BAT;.CMD;.VBS;.VBE;.JS;.JSE;.WSF;.WSH;.MSC;.CPL 才可以查找到
	cmdPath, err := exec.LookPath("shics")
	fmt.Println("err =", err)
	fmt.Println("cmdPath =", cmdPath)
}
