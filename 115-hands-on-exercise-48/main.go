package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mm := []string{"Miss", "Moneypenny", "I'm 008."}

	xmd := [][]string{jb, mm}

	fmt.Println(jb)
	fmt.Println(mm)
	fmt.Println(xmd)

	for _, v1 := range xmd {
		for _, v2 := range v1 {
			fmt.Printf("\"%v\" ", v2)
		}
		fmt.Println()
	}
}

/*

 "James", "Bond", "Shaken, not stirred"
 "Miss", "Moneypenny", "I'm 008."
*/
