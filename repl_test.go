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
			input:    " ",
			expected: []string{},
		},
		{
			input:    " hello ",
			expected: []string{"hello"},
		},
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "SoMe CrazY Text ",
			expected: []string{"some", "crazy", "text"},
		},
	}
	for _, c := range testCases {
		actual := CleanInput(c.input)
		for i := range c.expected {
			if actual[i] != c.expected[i] {
				t.Errorf("not equal: %v, %v", actual[i], c.expected[i])
				t.Fail()
			}
		}
	}
}
