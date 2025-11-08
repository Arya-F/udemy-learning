package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	const gs int = 100
	var counter int = 0

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("inner inner counter \t", counter)
			wg.Done()
		}()
		// fmt.Println("inner counter \t", counter)
	}
	wg.Wait()
	fmt.Println("counter:\t\t", counter)
}
