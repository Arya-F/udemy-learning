package main

import (
	"fmt"
	"time"
)

func main() {
	TimedFunction(MyFunction)
}

func TimedFunction(fn func()) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	fmt.Println("Elapsed time:", elapsed)
}

// Function to be wrapped
func MyFunction() {
	// time.Sleep(2 * time.Second) //Simulate some work
	for i := 0; i < 101; i++ {
		fmt.Println(i)
	}
	fmt.Println("MyFunction completed")
}
