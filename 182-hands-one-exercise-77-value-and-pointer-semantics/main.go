package main

import "fmt"

type person struct {
	first string
}

func main() {
	p1 := person{
		first: "James",
	}
	p2 := person{
		first: "Jenny",
	}
	fmt.Println(p1.first)

	p1 = changeName(p1, "Bond")
	fmt.Println(p1.first)

	fmt.Println(p2.first)
	changeNamePointer(&p2, "Monepenny")
	fmt.Println(p2.first)

}

func changeName(p person, s string) person {
	p.first = s
	return p
}

func changeNamePointer(p *person, s string) {
	p.first = s
}
