package main

import (
	"fmt"
)

func main() {
	jb := []string{"james", "bond", "martini", "chocolate"}
	jm := []string{"jenny", "moneypenny", "guiness", "wolverine tracks"}

	xss := [][]string{jb, jm}

	fmt.Println(jb)
	fmt.Println(jm)
	fmt.Println(xss)
	fmt.Println(xss[0][3])

	xi := []int{1, 2, 3}
	xi = append(xi, xi...)
	fmt.Println(xi)
}
