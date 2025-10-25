package main

import "fmt"

func main() {

	i := 0

	for {
		if i > 10 {
			break
		}
		fmt.Printf("Value of i is %v\n", i)
		i++
	}

}
