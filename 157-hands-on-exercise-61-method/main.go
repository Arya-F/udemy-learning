package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) Speak() {
	fmt.Printf("My name is %v, and I'm %v years old", p.first, p.age)
}

func main() {
	p1 := person{
		first: "James",
		age:   42,
	}

	p1.Speak()

}
