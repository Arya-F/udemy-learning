package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) {
	// _, err := w.Write([]byte(p.first))
	w.Write([]byte(p.first))
	// return err
}

func main() {
	p1 := person{
		first: "James",
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer f.Close()

	var b bytes.Buffer

	p1.writeOut(f)
	p1.writeOut(&b)

	fmt.Println(b.String())

}
