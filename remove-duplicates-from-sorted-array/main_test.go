package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		// {[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}
	for _, tt := range cases {
		actual := removeDuplicates(tt.nums)
		if actual != tt.expected {
			t.Errorf("removeDuplicates(%v): expected %d, but got %d", tt.nums, tt.expected, actual)
		}
	}
}
