package main

import "fmt"

func main() {
	bond := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Jane": 2,
		},
		favDrinks: []string{"Water", "Tea"},
	}
	moneypenny := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Jenny",
		friends: map[string]int{
			"Bond": 1,
		},
		favDrinks: []string{"Coffee", "Latte"},
	}

	fmt.Println(bond)
	fmt.Println(moneypenny)
}
