package main

import (
	"testing"
)

type testCase struct {
	input    string
	expected []string
}

func TestCleanInput(t *testing.T) {
	testCases := []testCase{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range testCases {
		actual := cleanInput(c.input)
		for i := range c.expected {
			if actual[i] != c.expected[i] {
				t.Errorf("not equal: %v, %v", actual[i], c.expected[i])
				t.Fail()
			}
		}
	}
}
