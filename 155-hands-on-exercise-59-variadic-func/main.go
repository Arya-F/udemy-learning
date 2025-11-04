package main

import (
	"fmt"
)

func main() {

	ix := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(foo(ix...))
	fmt.Println("---------")
	fmt.Println(bar(ix))

}

func foo(xi ...int) int {
	fmt.Println("Foo ran")
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	fmt.Println("Bar ran")
	total := 0
	for _, v := range xi {
		total += v
	}

	return total
}
