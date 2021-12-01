package messages_test

import (
	"testing"

	"github.com/tzarick/go-learning/getting-started/testing/messages"
)

func TestGreet(t *testing.T) {
	got := messages.Greet("gopher")
	expect := "Hello, gopher\n"
	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got %q", expect, got)
	}
}

func TestDepart(t *testing.T) {
	got := messages.Depart("gopher")
	expect := "Goodbye, gopher\n"
	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got %q", expect, got)
	}
}

// func TestFailureTypes(t *testing.T) {
// 	t.Error("Error - but doesn't stop the rest of the test from executing")
// 	t.Fatal("Fatal will fail the test and stop its execution")
// 	t.Error("errrrrror")
// }

func TestGreetTableDriven(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "gopher", expect: "Hello, gopher\n"},
		{input: "", expect: "Hello, \n"},
	}
	for _, s := range scenarios {
		got := messages.Greet(s.input)
		if got != s.expect {
			t.Errorf("Did not get expected result for input %v. Wanted %q, got %q", s.input, s.expect, got)
		}
	}
}
