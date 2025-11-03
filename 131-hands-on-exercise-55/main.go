package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {
	carA := vehicle{
		engine: engine{electric: false},
		make:   "Honda",
		model:  "Brio",
		doors:  4,
		color:  "Black",
	}
	carB := vehicle{
		engine: engine{electric: true},
		make:   "Lamborghini",
		model:  "Aventador",
		doors:  2,
		color:  "Blue",
	}

	fmt.Println(carA)
	fmt.Println(carB)

	fmt.Println(carA.electric, carA.make, carA.model, carA.doors, carA.color)
	fmt.Println(carB.electric, carB.make, carB.model, carB.doors, carB.color)

}
