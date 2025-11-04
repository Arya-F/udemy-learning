package main

import "fmt"

func main() {

	x := doMath(42, 69, add)
	fmt.Println(x)

	y := doMath(42, 69, subtract)
	fmt.Println(y)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
