package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	cases := map[string]struct {
		arg      string
		expected bool
	}{
		"s is empty":                              {arg: "", expected: false},
		"input `()` expect true":                  {arg: "()", expected: true},
		"input `()[]{}` expect true":              {arg: "()[]{}", expected: true},
		"input `(]` expect false":                 {arg: "(]", expected: false},
		"input `{ [ { } ] } ( ) { }` expect true": {arg: "{[{}]}(){}", expected: true},
		"false": {arg: "{[{}}){}", expected: false},
	}
	for desc, tt := range cases {
		t.Run(desc, func(t *testing.T) {
			actual := isValid(tt.arg)
			if actual != tt.expected {
				t.Errorf("expected %v, but got %v", tt.expected, actual)
			}
		})
	}
}

func TestIsValidV2(t *testing.T) {
	cases := map[string]struct {
		arg      string
		expected bool
	}{
		"s is empty":                              {arg: "", expected: false},
		"input `()` expect true":                  {arg: "()", expected: true},
		"input `()[]{}` expect true":              {arg: "()[]{}", expected: true},
		"input `(]` expect false":                 {arg: "(]", expected: false},
		"input `{ [ { } ] } ( ) { }` expect true": {arg: "{[{}]}(){}", expected: true},
		"false": {arg: "{[{}}){}", expected: false},
	}
	for desc, tt := range cases {
		t.Run(desc, func(t *testing.T) {
			actual := isValidV2(tt.arg)
			if actual != tt.expected {
				t.Errorf("expected %v, but got %v", tt.expected, actual)
			}
		})
	}
}
