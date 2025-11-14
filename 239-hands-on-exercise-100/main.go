package main

import "fmt"

func main() {

}

func multiply(a, b int) int {
	return a * b
}

func division(a, b int) (float64, error) {
	if a == 0 || b == 0 {
		return 0, fmt.Errorf("error, divison by zero")
	}
	return float64(a) / float64(b), nil
}
