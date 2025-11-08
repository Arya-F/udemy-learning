package main

import "fmt"

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println(p.first, "is speaking")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p1 := person{"James"}
	p2 := person{"Jenny"}

	p3 := &p2

	//saySomething(p1)// won't work
	saySomething(&p2) // will work
	p1.speak()        // will work
	saySomething(p3)  // will work

}
