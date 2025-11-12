package main

import "fmt"

func main() {
	c1 := make(chan int)
	go func() {
		for i := range 100 {
			c1 <- i
		}
		close(c1)
	}()
	for v := range c1 {
		fmt.Println("value of c1: \t", v)
	}
	fmt.Println("about to exit")

}
