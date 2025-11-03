package main

import "fmt"

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {
	p1 := person{
		first:    "James",
		last:     "Bond",
		iceCream: []string{"Chocolate", "Black Forest", "Vanilla"},
	}

	p2 := person{
		first:    "Jenny",
		last:     "Moneypenny",
		iceCream: []string{"Strawberry", "Red Velvet", "Caramel"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last)
	for i, v := range p1.iceCream {
		fmt.Println(i, v)
	}
	fmt.Println(p2.first, p2.last)
	for i, v := range p2.iceCream {
		fmt.Println(i, v)
	}
}
