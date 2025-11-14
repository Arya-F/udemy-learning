package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintln("there is an error:", ce.info)
}

func main() {
	c1 := customErr{"404 Found"}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
