package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) calcArea() float64 {
	return s.length * s.width
}

func (c circle) calcArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shapes interface {
	calcArea() float64
}

func info(s shapes) {
	fmt.Println(s.calcArea())
}

func main() {
	x := square{
		length: 10,
		width:  20,
	}

	y := circle{
		radius: 10,
	}

	fmt.Println(x.calcArea())
	fmt.Println(y.calcArea())
	fmt.Println("------------")
	info(x)
	info(y)

}

/*
 ● create a type SQUARE
 ○ length float64
 ○ width float64
 ● create a type CIRCLE
 ○ radius float64
 ● attach a method to each that calculates AREA and returns it
 ○ circle area= π r 2
 ■ math.Pi
 ■ math.Pow
 ○ square area = L * W
 ● create a type SHAPE that defines an interface as anything that has the AREA method
 ● create a func INFO which takes type shape and then prints the area
 ● create a value of type square
 ● create a value of type circle
 ● usefunc info to print the area of square
 ● usefunc info to print the area of circl

*/
