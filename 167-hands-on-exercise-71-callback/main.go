package main

import (
	"fmt"
)

func main() {
	x := printSquare(square, 5)

	fmt.Println(x)
}

func square(i int) int {
	return i * i
}

func printSquare(f func(int) int, i int) string {
	return fmt.Sprintf("Area of %v is %v", i, f(i))
}
