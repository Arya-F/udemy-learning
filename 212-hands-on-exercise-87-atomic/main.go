package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	const gs int = 100
	var counter int64 = 0

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter : \t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		// fmt.Println("inner counter \t", counter)
	}
	wg.Wait()
	fmt.Println("counter:\t\t", counter)
}
