package main

import (
	"fmt"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

func main() {
	// 要获取的环境变量名
	name := "PATH"

	// 加载 kernel32.dll 库
	kernel32, err := syscall.LoadDLL("kernel32.dll")
	if err != nil {
		fmt.Println("无法加载 kernel32.dll:", err)
		return
	}

	// 获取 GetEnvironmentVariableW 函数的地址
	getEnvironmentVariableW, err := kernel32.FindProc("GetEnvironmentVariableW")
	if err != nil {
		fmt.Println("无法找到 GetEnvironmentVariableW 函数:", err)
		return
	}

	// 创建一个缓冲区来存储环境变量值
	buffer := make([]uint16, syscall.MAX_PATH+1)

	// 调用 GetEnvironmentVariableW 函数
	ret, _, err := syscall.Syscall(getEnvironmentVariableW.Addr(), 3,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(len(buffer)))

	// 检查调用结果和错误
	if ret == 0 {
		fmt.Println("无法获取环境变量值:", err)
		return
	}

	// 将缓冲区转换为字符串并打印环境变量值
	value := string(utf16.Decode(buffer[:ret-1]))
	fmt.Println("环境变量值:", value)
}
