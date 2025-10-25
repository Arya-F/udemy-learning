package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var x int = rand.Intn(500)
	fmt.Printf("variable x value is %v with type of %T\n", x, x)

	// fmt.Println("if Statement")
	// if x <= 100 {
	// 	fmt.Println("between 0 and 100")
	// } else if x >= 101 && x <= 200 {
	// 	fmt.Println("between 101 and 200")
	// } else if x >= 201 && x <= 250 {
	// 	fmt.Println("between 201 and 250")
	// } else {
	// 	fmt.Println("More than 250")
	// }
	fmt.Println("Switch Statement")
	switch {
	case x <= 100:
		fmt.Println("between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("between 201 and 250")
	case x >= 201 && x <= 250:
		fmt.Println("between 201 and 250")
	default:
		fmt.Println("More than 250")
	}
}
