package main

import "testing"

func TestBalancedSums(t *testing.T) {
	cases := []struct {
		desc     string
		arr      []int32
		expected string
	}{
		{desc: "", arr: []int32{2, 0, 0, 0}, expected: "YES"},
		{desc: "", arr: []int32{1, 2, 3}, expected: "NO"},
		{desc: "", arr: []int32{1, 2, 3, 3}, expected: "YES"},
	}
	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			actual := BalancedSums(tt.arr)
			if actual != tt.expected {
				t.Errorf("expected %s, but got %s", tt.expected, actual)
			}
		})
	}
}
