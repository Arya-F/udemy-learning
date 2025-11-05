package main

import (
	"fmt"
	"math"
)

func main() {
	f := calcPow(2)
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func calcPow(x float64) func() float64 {
	var i float64 = 0
	return func() float64 {
		i++
		return math.Pow(x, i)
	}
}
