package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("=================")
	fmt.Println(xi[:5])
	fmt.Println("=================")
	fmt.Println(xi[5:])
	fmt.Println("=================")
	fmt.Println(xi[2:7])
	fmt.Println("=================")
	fmt.Println(xi[1:6])
}
