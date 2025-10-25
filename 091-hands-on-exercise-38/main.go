package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 100; i++ {
		fmt.Printf("No. %v \t", i)
		if x := rand.Intn(5); x == 3 {
			fmt.Println("X is 3", x)
		} else {
			fmt.Println("X is not 3", x)
		}
	}
}
