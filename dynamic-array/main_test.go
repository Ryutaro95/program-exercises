package main

import (
	"reflect"
	"testing"
)

func TestDynacmicArray(t *testing.T) {
	cases := []struct {
		desc     string
		n        int32
		queries  [][]int32
		expected []int32
	}{
		{
			desc:     "",
			n:        2,
			queries:  [][]int32{{1, 0, 5}, {1, 1, 7}, {1, 0, 3}, {2, 1, 0}, {2, 1, 1}},
			expected: []int32{7, 3},
		},
	}
	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			actual := dynamicArray(tt.n, tt.queries)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, actual)
			}
		})
	}
}
