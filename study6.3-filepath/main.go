package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	res := filepath.Base("/path/to/file.txt")
	fmt.Println(res)
}
