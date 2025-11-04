package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := readFile("text.txt")

	if err != nil {
		log.Fatalf("error in main %s", err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))
}

func readFile(filename string) ([]byte, error) {
	xb, err := os.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("error in readFile %s", err)
	}
	return xb, nil
}
