package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "	hello	world		",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Luke I am your father",
			expected: []string{"luke", "i", "am", "your", "father"},
		},
		{
			input:    ".this, is. a ,test,",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected %v, got %v", expectedWord, word)
			}
		}
	}
}
