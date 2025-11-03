package main

import "fmt"

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {

	peoples := make(map[string]person)

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

	fmt.Println("----------------------------------------------")

	peoples[p1.last] = p1
	peoples[p2.last] = p2

	fmt.Println(peoples)
	fmt.Println("----------------------------------------------")

	fmt.Println(peoples["Bond"].first, peoples["Bond"].last, peoples["Bond"].iceCream)
	fmt.Println(peoples["Moneypenny"].first, peoples["Moneypenny"].last, peoples["Moneypenny"].iceCream)
	fmt.Println("----------------------------------------------")

	for _, v := range peoples {
		fmt.Printf("%v favorites are \n", v.first)
		for _, v2 := range v.iceCream {
			fmt.Println(v2)
		}
	}

}
