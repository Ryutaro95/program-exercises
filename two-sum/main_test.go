package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := map[string]struct {
		nums     []int
		expected []int
		target   int
	}{
		"a": {target: 9, nums: []int{2, 7, 11, 15}, expected: []int{0, 1}},
		"b": {target: 6, nums: []int{3, 2, 3}, expected: []int{0, 2}},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			actual := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, actual)
			}
		})
	}
}
