package main

import "testing"

func TestStringXOR(t *testing.T) {
	testCases := []struct {
		desc, s, t, expected string
	}{
		{
			desc:     "",
			s:        "",
			t:        "",
			expected: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := stringsXOR(tC.s, tC.t)
			if actual != tC.expected {
				t.Errorf("expected %s, but got %s.", tC.expected, actual)
			}
		})
	}
}
