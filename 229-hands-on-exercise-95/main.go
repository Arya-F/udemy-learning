// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	c1 := make(chan int)

// 	go populate(c1)

// 	for v := range c1 {
// 		fmt.Println("value from c1: \t", v)
// 	}
// }

// func populate(c chan int) {
// 	var wg sync.WaitGroup
// 	const goroutines = 10
// 	wg.Add(goroutines)

// 	for i := 0; i < goroutines; i++ {
// 		go func() {
// 			for j := 0; j < goroutines; j++ {
// 				c <- j
// 			}
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	close(c)
// }

package main

import "fmt"

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for j := range 10 {
				c <- j
			}
		}()
	}

	for i := range 100 {
		fmt.Println(i, <-c)
	}
	fmt.Println("about to exit")
}
