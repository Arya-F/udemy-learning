package main

import (
	"fmt"
)

func main() {

	fmt.Println("First")
	defer fmt.Println("second")
	fmt.Println("Third")
	defer fmt.Println("Fourth")
	defer fmt.Println("Fifth")

}
