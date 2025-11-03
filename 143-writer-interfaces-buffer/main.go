package main

import (
	"bytes"
	"fmt"
)

// type person struct {
// 	first string
// }

// func (p person) writeOut(w io.Writer) error {
// 	_, err := w.Write([]byte(p.first))
// 	return err
// }

func main() {
	// f, err := os.Create("output.txt")
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }

	// defer f.Close()

	// s := []byte("Hello sekai")

	// _, err = f.Write(s)

	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }

	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	b.WriteString("Sekai")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("Ohayo World")
	fmt.Println(b.String())

	b.Write([]byte("Good Morning Sekai"))
	fmt.Println(b.String())

}
