package main

import "fmt"

func main() {
	foo()

	x := func() {
		fmt.Println("Anonymous func")
	}

	y := func(s string) {
		fmt.Println("Anonymous func with argument", s)
	}

	x()
	y("James")

}

func foo() {
	fmt.Println("Foo")
}
