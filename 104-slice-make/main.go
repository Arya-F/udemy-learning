package main

import (
	"fmt"
)

func main() {
	xi := make([]int, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("-------------------")
	xi[5] = 80085
	xi[4] = 69420
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("-------------------")
	xi = append(xi[:4], xi[6:]...)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("-------------------")
	xi = append(xi, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
}
