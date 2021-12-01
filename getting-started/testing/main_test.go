package main_test

import "testing"

// test std lib
func TestAddition(t *testing.T) {
	got := 2 + 2
	expected := 4
	if got != expected {
		t.Errorf("Did not get expected result. Got '%v', wanted '%v'", got, expected)
	}
}

func TestSubtraction(t *testing.T) {
	got := 10 - 5
	expected := 5
	if got != expected {
		t.Errorf("Did not get expected result. Got '%v', wanted '%v'", got, expected)
	}
}
