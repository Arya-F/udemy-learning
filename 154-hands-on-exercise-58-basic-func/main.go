package main

import (
	"fmt"
)

func main() {
	x := foo()
	y, z := bar()

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func foo() int {
	fmt.Println("Foo ran")
	return 42
}

func bar() (string, int) {
	return "Bar ran", 69
}
