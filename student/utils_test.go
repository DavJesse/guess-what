package main

import "testing"

func TestAtoi(t *testing.T) {
	test := "TestAtoi"
	subject := "123"
	gotInt, gotErr := atoi(subject) // Extract output for testing
	expectedInt := 123

	if gotInt != expectedInt || gotErr != "" {
		t.Errorf("Got: %v", gotInt)
		t.Errorf("Expected: %v", expectedInt)
		t.Errorf("%s Failed!", test)
	}
}

func TestIsNumeric(t *testing.T) {
	test := "TestIsNumeric"
	subject := "3MIN3M"
	got := isNumeric(subject) // Extract output for testing
	expected := false

	if got != expected {
		t.Errorf("Got: %t", got)
		t.Errorf("Expected: %t", expected)
		t.Errorf("%s Failed!", test)
	}
}

// func Test(t *testing.T) {
// 	test := "Test"
// 	subject :=
// 	got := // Extract output for testing
// 	expected :=

// 	if got != expected {
// 		t.Errorf("Got: %v", got)
// 		t.Errorf("Expected: %v", expected)
// 		t.Errorf("%s Failed!", test)
// 	}
// }
