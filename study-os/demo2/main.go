package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("dir")
	cmd.Stdout = os.Stdout
	// Run 和 Start只能用一个
	// Run starts the specified command and waits for it to complete.
	_ = cmd.Run()

	// Start starts the specified command but does not wait for it to complete.
	// _ = cmd.Start()
	// _ = cmd.Wait()
}
