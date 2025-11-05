package main

import "fmt"

func main() {
	x := bar()
	fmt.Println(x(5))
}

func bar() func(j int) int {
	fmt.Println("bar ran")
	return func(j int) int {
		fmt.Println("return func ran")
		return 43 + j

	}
}
