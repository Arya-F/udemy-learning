package main

import "fmt"

func main() {
	xm := map[string][]string{
		"bond_james":       {"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"intelligence", "literature", "computer science"},
		"no_dr":            {"cats", "ice cream", "sunsets"},
	}

	fmt.Println(xm)
	fmt.Println("-----------------------")

	for k, v := range xm {
		fmt.Println(k, v)
		for i, j := range v {
			fmt.Println(i, j)
		}
	}
}
