package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	cases := map[string]struct {
		expected string
		strs     []string
	}{
		"expect fl":    {strs: []string{"flower", "flow", "flight"}, expected: "fl"},
		"":             {strs: []string{"flight"}, expected: "flight"},
		"expect empty": {strs: []string{"dog", "racecar", "car"}, expected: ""},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			actual := longestCommonPrefix(tt.strs)
			if actual != tt.expected {
				t.Errorf("expected %s, but got %s", tt.expected, actual)
			}
		})
	}
}
