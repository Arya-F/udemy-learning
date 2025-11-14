package main

import (
	"fmt"
	dog "udemylearning/245-hands-on-exercise-101/package"
)

func main() {

	hAge := 10
	dAge := dog.Years(hAge)
	fmt.Println(dAge)
}
