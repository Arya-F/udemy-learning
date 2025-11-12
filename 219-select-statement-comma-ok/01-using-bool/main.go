package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go send(even, odd, quit)

	// receive
	receive(even, odd, quit)

	fmt.Println("About to quit")

}

func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("From even channel: \t", v)
		case v := <-odd:
			fmt.Println("From odd channel: \t", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok: \t", i, ok)
				return
			} else {
				fmt.Println("from comma ok: \t", i)
			}
		}
	}
}
func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}
