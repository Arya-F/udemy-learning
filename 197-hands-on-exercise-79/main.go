package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	// fmt.Println(s)
	var people []person

	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, person := range people {
		fmt.Printf("No. %v, %v %v, age %v is saying\n", i, person.First, person.Last, person.Age)
		for j, v2 := range person.Sayings {
			fmt.Printf("\t\t%v %v \n", j, v2)
		}
	}

}
