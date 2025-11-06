package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPwd1 := `password124`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPwd1))
	if err != nil {
		// log.Fatalf("Error: %v", err)
		fmt.Println("error")
		return
	}
	fmt.Println("You're logged in")
}
