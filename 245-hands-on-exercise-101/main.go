package main

import (
	"fmt"

	"github.com/Arya-F/udemy-learning/245-hands-on-exercise-101/dog"
)

func main() {

	hAge := 10
	dAge := dog.Years(hAge)
	fmt.Println(dAge)
}
