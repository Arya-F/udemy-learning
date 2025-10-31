package main

import "fmt"

func main() {
	xm := make(map[string][]string)
	xm["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	xm["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	xm["no_dr"] = []string{"cats", "ice cream", "sunsets"}

	fmt.Println(xm)
	fmt.Println("-----------------------")

	for k, v := range xm {
		fmt.Println(k, v)
		for i, j := range v {
			fmt.Println(i, j)
		}
	}

	fmt.Println("-----------------------")

	xm["fleming_ian"] = []string{"steaks", "cigars", "espionage"}
	for k, v := range xm {
		fmt.Println(k, v)
		for i, j := range v {
			fmt.Println(i, j)
		}
	}
}
