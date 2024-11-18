package main

import (
	"fmt"
	"os"
)

func main() {
	// read a file
	bytes, _ := os.ReadFile("./example/00.lang")
	source := string(bytes)

	fmt.Printf("code: %s\n", source)
}
