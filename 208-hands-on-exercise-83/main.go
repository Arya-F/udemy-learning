package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	fmt.Println("Begin CPU", runtime.NumCPU())
	fmt.Println("Begin gs", runtime.NumGoroutine())

	go countToX1(10)
	go countToX2(10)
	go countToX3(10)

	fmt.Println("Mid CPU", runtime.NumCPU())
	fmt.Println("Mid gs", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("End CPU", runtime.NumCPU())
	fmt.Println("End gs", runtime.NumGoroutine())
}

func countToX1(x int) {
	for i := 0; i < x; i++ {
		fmt.Println("First No.", i)
	}
	wg.Done()
}
func countToX2(x int) {
	for i := 0; i < x; i++ {
		fmt.Println("Second No.", i)
	}
	wg.Done()
}
func countToX3(x int) {
	for i := 0; i < x; i++ {
		fmt.Println("Third No.", i)
	}
	wg.Done()
}
