package main

import "fmt"

func intDelta(i *int) {
	*i = 69
}

func sliceDelta(ii []int) {
	ii[0] = 99
}

func mapDelta(md map[string]int, s string) {
	md[s] = 32
}

func main() {
	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)

	m := map[string]int{
		"james": 33,
	}
	fmt.Println(m["james"])
	mapDelta(m, "james")
	fmt.Println(m["james"])

}
