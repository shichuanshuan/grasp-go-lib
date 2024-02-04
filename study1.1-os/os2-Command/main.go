package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func UseCommand() {
	// cmd /C 用于执行完命令后关闭命令提示符窗口
	cmd := exec.Command("cmd", "/C", "echo", "Hello world")

	//cmd := exec.Command("python", "shics.py")
	//cmd := exec.Command("D:\\python\\python27\\python", "shics.py")

	// 1
	// 执行命令并获取输出结果
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("命令执行失败:", err)
		return
	}

	// 输出命令的输出结果
	fmt.Println(string(output))

	// 2
	//cmd.Stdout = os.Stdout
	// Run 和 Start只能用一个
	// Run starts the specified command and waits for it to complete.
	//_ = cmd.Run()

	// 3
	// Start starts the specified command but does not wait for it to complete.
	//_ = cmd.Start()
	//_ = cmd.Wait()
}

func UseCommandContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	//if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
	//cmd := exec.CommandContext(ctx, "sleep", "5")
	cmd := exec.CommandContext(ctx, "timeout", "/t", "5", "nobreak")
	if err := cmd.Run(); err != nil {
		// This will fail after 100 milliseconds. The 5 second sleep
		// will be interrupted.
		fmt.Println("failed...", err)
	}

	fmt.Println("success...")
}

func main() {
	//UseCommand()

	UseCommandContext()
}
