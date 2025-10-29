package main

import "fmt"

func main() {
	arrayInt := [5]int{}

	for i := 0; i < 5; i++ {
		arrayInt[i] = i + 2
	}

	fmt.Println(arrayInt)
	fmt.Println(len(arrayInt))
	fmt.Println(cap(arrayInt))
	fmt.Println("=================")

	for i, v := range arrayInt {
		fmt.Printf("Index : %v \t Value: %v \t Type: %T\n", i, v, v)
	}
}
