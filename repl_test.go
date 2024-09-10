package main

import (
	"testing"
)

func TestRepl(t *testing.T) {
	tests := []struct{
		name string
		input string
		expected []string
	}{
		{
			name: "empty input",
			input: "",
			expected: []string{},
		},
		{
			name: "strip whitespace",
			input: "  hello ",
			expected: []string{"hello"},
		},
		{
			name: "strip whitespace and segment words",
			input: "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			name: "strip whitespace, segment words and lowercase",
			input: "  HellO, World!  ",
			expected: []string{"hello,", "world!"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := cleanInput(tc.input)

			if len(actual) != len(tc.expected) {
				t.Errorf(
					"FAIL: length of actual %v: '%v' does not match length of expected %v: '%v'",
					actual, len(actual), tc.expected, len(tc.expected))
			}

			for j, word := range actual {
				expectedWord := tc.expected[j]
				if word != expectedWord {
					t.Errorf("Test %v - '%s' FAIL: cleanInput(%v) == %v, expected %v",
					i, tc.name, tc.input, actual, tc.expected)
				}
			}
		})
	}
}