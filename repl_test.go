package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " test the repl ",
			expected: []string{"test", "the", "repl"},
		},
		{
			input:    " is this a clean string  ",
			expected: []string{"is", "this", "a", "clean", "string"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

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
