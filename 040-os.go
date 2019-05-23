package main

import (
	"fmt"
	"os"
)

func main() {
	err1 := os.Remove("/tmp/test")
	if err1 != nil {
		fmt.Printf("ERROR: %s\n", err1.Error())
	}
	err := os.Mkdir("/tmp/test", 0777)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
	tmpDir := os.TempDir()
	fmt.Println(tmpDir)

	var _ string = ""
}
