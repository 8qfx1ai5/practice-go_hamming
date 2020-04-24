package main

import "testing"

// TestHamming run a test on the hamming function
// TODO: add more test-cases
func TestHamming(t *testing.T) {
	var input1 string
	var input2 string
	var result int
	var expected int

	input1 = "foo"
	input2 = "ba3"
	expected = 3
	result = hamming(input1, input2)
	if result != expected {
		t.Errorf("unequal hamming strings, give %s %s, expected %d, result %d", input1, input1, expected, result)
	}
}
