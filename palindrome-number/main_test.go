package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := map[string]struct {
		x        int
		expected bool
	}{
		"a": {x: 121, expected: true},
		"b": {x: -121, expected: false},
		"c": {x: 10, expected: false},
		"d": {x: 1234567654321, expected: true},
		"e": {x: 0, expected: true},
	}
	for desc, tC := range testCases {
		t.Run(desc, func(t *testing.T) {
			actual := isPalindrome(tC.x)
			if actual != tC.expected {
				t.Errorf("expected %t, but got %t", tC.expected, actual)
			}
		})
	}
}
