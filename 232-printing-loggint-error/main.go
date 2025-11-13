package main

import (
	"log"
	"os"
)

func main() {
	f, _ := os.Create("log.txta")
	defer f.Close()

	log.SetOutput(f)

	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("Error: \t", err)
		log.Println("Error: \t", err)
		// log.Fatalln("Error: \t", err)
		// panic(err)
	}

}
