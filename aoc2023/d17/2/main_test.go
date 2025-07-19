package main

import "testing"

func TestCalc(t *testing.T) {
	r := Calc("0.txt")
	if r != 102 {
		t.Fatalf("Expected 102, got %v\n", r)
	}

}
