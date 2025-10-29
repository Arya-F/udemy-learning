package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("=================")

	for i, v := range xi {
		fmt.Printf("Index : %v \t Value: %v \t Type: %T\n", i, v, v)
	}
}
