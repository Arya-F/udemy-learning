package main

import "testing"

func TestMultipy(t *testing.T) {
	got := multiply(5, 9)
	if got != 45 {
		t.Errorf("multiply(5,9) want:45\tgot:%v", got)
	}
}

func TestDivision(t *testing.T) {
	got, err := division(10, 5)
	if err != nil {
		t.Errorf("division(10,5)=%v, want 2", got)
	}
}
