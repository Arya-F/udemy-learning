package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 9)
	if total != 14 {
		t.Errorf("Wrong! got %d, need %d", total, 14)
	}
}
