package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Test")
	fmt.Fprintln(os.Stdout, "Test2")
	io.WriteString(os.Stdout, "Test3")
}
