package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("Nice")
	}

	func() {
		fmt.Println("Nice2")
	}()

	y := func(s string) {
		fmt.Println("Her name is", s)
	}

	x()

	func(s string) {
		fmt.Println("my name is", s)
	}("James")

	y("Jenny")
}
