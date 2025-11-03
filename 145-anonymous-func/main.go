package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func")
	}()

	func(s string) {
		fmt.Println("Anonymous func with argument", s)
	}("Jason")
}

func foo() {
	fmt.Println("Foo")
}
