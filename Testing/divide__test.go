package main

import "testing"

func TestDivideByZero(t *testing.T) {
	_, err := divide(1, 0)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

func TestDivide(t *testing.T) {
	result, _ := divide(1, 1)
	expected := 1
	if result != expected {
		t.Errorf("Expected the result to be %d but got %d", expected, result)
	}
}
