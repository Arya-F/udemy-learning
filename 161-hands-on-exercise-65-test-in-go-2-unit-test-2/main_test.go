package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(5, 9)
	if total != 14 {
		t.Errorf("Wrong! got %d, need %d", total, 14)
	}
}

func TestSubtract(t *testing.T) {
	total := subtract(5, 9)
	if total != -4 {
		t.Errorf("Wrong! got %d, need %d", total, -4)
	}
}

func TestDoMath(t *testing.T) {
	total := doMath(5, 9, add)
	if total != 14 {
		t.Errorf("Wrong! got %d, need %d", total, 14)
	}
}
func TestMultiply(t *testing.T) {
	total := doMath(5, 9, multiply)
	if total != 45 {
		t.Errorf("Wrong! got %d, need %d", total, 45)
	}
}
