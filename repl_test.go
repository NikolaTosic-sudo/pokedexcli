package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    " hello world ",
			expected: "hello",
		},
		{
			input:    " test the repl ",
			expected: "test",
		},
		{
			input:    " is this a clean string  ",
			expected: "is",
		},
	}

	for _, c := range cases {
		actual, _ := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("expected length %d, got %d", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("expected length %d, got %d", len(c.expected), len(actual))
			}
		}
	}
}
