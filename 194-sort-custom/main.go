package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByName []Person

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Less(i, j int) bool {
	return a[i].First < a[j].First
}

// type ByAge []Person

// func (a ByAge) Len() int {
// 	return len(a)
// }

// func (a ByAge) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

// func (a ByAge) Less(i, j int) bool {
// 	return a[i].Age < a[j].Age
// }

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

}

// package main

// import (
// 	"fmt"
// )

// // Define a custom type implementing the sort.Interface methods.

// type Person struct {
// 	First string
// 	Age   int
// }

// type ByName []Person

// func (a ByName) Len() int {
// 	return len(a)
// }

// func (a ByName) Less(i, j int) bool {
// 	return a[i].First < a[j].First
// }

// func (a ByName) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

// // SelectionSort simulates what sort.Sort does internally:
// // It repeatedly uses Len(), Less(), and Swap() to sort the slice.
// func SelectionSort(data ByName) {
// 	n := data.Len()
// 	for i := 0; i < n-1; i++ {
// 		minIndex := i
// 		for j := i + 1; j < n; j++ {
// 			if data.Less(j, minIndex) {
// 				minIndex = j
// 			}
// 		}
// 		if minIndex != i {
// 			data.Swap(i, minIndex)
// 		}
// 	}
// }

// func main() {
// 	people := ByName{
// 		{"James", 32},
// 		{"Moneypenny", 27},
// 		{"Q", 64},
// 		{"M", 56},
// 	}

// 	fmt.Println("Before:", people)
// 	SelectionSort(people) // Replace with sort.Sort to do the same with Go's built-in implementation
// 	fmt.Println("After:", people)
// }
