package main

import "testing"

func TestGridChallenge(t *testing.T) {
	cases := []struct {
		desc     string
		grid     []string
		expected string
	}{
		{
			desc:     "test 1",
			grid:     []string{"abc"},
			expected: "YES",
		},
		{
			desc:     "test 1",
			grid:     []string{"abc", "lmp", "qrt"},
			expected: "YES",
		},
		{
			desc:     "test 2",
			grid:     []string{"mpxz", "abcd", "wlmf"},
			expected: "NO",
		},
		{
			desc:     "test 3",
			grid:     []string{"abc", "hjk", "mpq", "rtv"},
			expected: "YES",
		},
		{
			desc:     "test 4",
			grid:     []string{"kc", "iu"},
			expected: "YES",
		},
		{
			desc:     "test 5",
			grid:     []string{"uxf", "vof", "hmp"},
			expected: "NO",
		},
	}
	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			actual := gridChallenge(tt.grid)
			if actual != tt.expected {
				t.Errorf("expected %s, but got %s", tt.expected, actual)
			}
		})
	}
}
