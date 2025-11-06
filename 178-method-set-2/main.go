package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println(d.first, "Is walking")
}

func (d *dog) run() {
	d.first = "Bond"
	fmt.Println(d.first, "Is running")
}

type young interface {
	walk()
	run()
}

func youngRun(y young) {
	y.run()
}

func main() {
	d1 := dog{"James"}
	d2 := &dog{"Jenny"}

	fmt.Println(d1)
	fmt.Println(*d2)

	d1.walk()
	d1.run()
	// youngRun(d1)

	d2.walk()
	d2.run()
	youngRun(d2)

	fmt.Println(d1)
	fmt.Println(*d2)

}
