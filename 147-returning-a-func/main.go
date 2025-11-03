package main

import "fmt"

func main() {
	foo()

	x := foo()
	fmt.Println(x)

	y := bar()

	fmt.Println(y())

}

func foo() int {
	return 42
}

func bar() func() int {
	return func() int {
		return 43
	}
}
