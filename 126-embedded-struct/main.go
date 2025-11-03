package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	// p1 := person{
	// 	first: "James",
	// 	last:  "Bond",
	// 	age:   42,
	// }
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   42,
		},
		ltk: true,
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", sa1, sa1)
	fmt.Printf("%T \t %#v\n", p2, p2)

	fmt.Println(sa1.person.age)

}
