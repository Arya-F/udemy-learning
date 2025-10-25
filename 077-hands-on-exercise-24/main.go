package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var x int = rand.Intn(500)
	fmt.Printf("variable x value is %v with type of %T\n", x, x)

	if x <= 100 {
		fmt.Println("between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("between 101 and 200")
	} else if x >= 201 && x <= 250 {
		fmt.Println("between 201 and 250")
	} else {
		fmt.Println("More than 250")
	}
}
