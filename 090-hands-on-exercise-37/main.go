package main

import "fmt"

func main() {

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	age := m["James"]
	fmt.Println(age)

	// age2 := m["Q"]
	// fmt.Println(age2)

	// for k, v := range m {
	// 	fmt.Printf("Key \"%v\" value is \"%v\"\n", k, v)
	// }

	if age2, ok := m["James"]; ok {
		fmt.Println(age2, !ok)
	} else {
		fmt.Println("key not found")
	}

}
