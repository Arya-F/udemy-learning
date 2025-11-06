package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Jenny","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)
	fmt.Printf("%T \n", s)
	fmt.Printf("%T \n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	fmt.Println(people)

	for i, v := range people {
		fmt.Printf("No. %v \t %v %v, %v\n", i, v.First, v.Last, v.Age)
	}

	var x int
	x, err = fmt.Println("123456789")
	fmt.Println(strconv.Itoa(x), err)
}
