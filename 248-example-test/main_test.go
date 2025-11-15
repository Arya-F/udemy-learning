package main

import (
	"testing"
)

func TestMultipy(t *testing.T) {
	type tests struct {
		num1   int
		num2   int
		answer int
	}

	test := []tests{
		{num1: 5, num2: 6, answer: 30},
		{num1: 7, num2: 7, answer: 49},
		{num1: 8, num2: 10, answer: 80},
	}

	for _, v := range test {

		got := multiply(v.num1, v.num2)
		if got != v.answer {
			t.Errorf("multiply(5,9) want:45\tgot:%v", got)
		}
	}
}

func TestDivision(t *testing.T) {
	type testsFloat struct {
		num1   int
		num2   int
		answer float64
	}

	test := []testsFloat{
		{num1: 30, num2: 6, answer: 5},
		{num1: 7, num2: 7, answer: 1},
		{num1: 80, num2: 10, answer: 8},
	}

	for _, v := range test {

		got, err := division(v.num1, v.num2)
		if err != nil {
			t.Errorf("division(10,5)=%v, want 2", got)
		}
		if got != v.answer {
			t.Errorf("Want: %v, Got: %v", v.answer, got)
		}
	}
}

// func ExampleMultiply() {
// 	fmt.Println(multiply(1, 2))
// 	//Output:
// 	//2
// }
