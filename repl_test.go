package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "helLo World",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "  helLo World  ",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The length are not equal: %v vs %v\n", len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := c.expected[i]
			if actualWord != expectedWord {
				fmt.Printf("Expect %v, but got %v\n", expectedWord, actualWord)
			}
		}
	}

}
