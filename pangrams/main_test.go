package main

import "testing"

func TestPangrams(t *testing.T) {
	cases := []struct {
		description string
		arg         string
		expected    string
	}{
		{"expect pangram because contains all alphabets", "We promptly judged antique ivory buckles for the next prize", "pangram"},
		{"expect not pangram because does not contain 'x'", "We promptly judged antique ivory buckles for the prize", "not pangram"},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			actual := pangrams(tt.arg)
			if actual != tt.expected {
				t.Errorf("expected %s, but got %s.", tt.expected, actual)
			}
		})
	}
}
