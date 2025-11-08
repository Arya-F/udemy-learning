package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	const gs int = 100
	var counter int = 0

	var mu sync.Mutex

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			// runtime.Gosched()
			v++
			counter = v
			fmt.Println("inner inner counter \t", counter)
			mu.Unlock()
			wg.Done()
		}()
		// fmt.Println("inner counter \t", counter)
	}
	wg.Wait()
	fmt.Println("counter:\t\t", counter)
}
