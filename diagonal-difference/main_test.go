package main

import "testing"

func TestDaigonalDifference(t *testing.T) {
	cases := []struct {
		description string
		arg         [][]int32
		expected    int32
	}{
		{
			"",
			[][]int32{
				{11, 2, 4},
				{4, 5, 6},
				{10, 8, -12},
			},
			15,
		},
		{
			"",
			[][]int32{
				{6, 6, 7, -10, 9, -3, 8, 9, -1},
				{9, 7, -10, 6, 4, 1, 6, 1, 1},
				{-1, -2, 4, -6, 1, -4, -6, 3, 9},
				{-8, 7, 6, -1, -6, -6, 6, -7, 2},
				{-10, -4, 9, 1, -7, 8, -5, 3, -5},
				{-8, -3, -4, 2, -3, 7, -5, 1, -5},
				{-2, -7, -4, 8, 3, -1, 8, 2, 3},
				{-3, 4, 6, -7, -7, -8, -3, 9, -6},
				{-2, 0, 5, 4, 4, 4, -3, 3, 0},
			},
			52,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			actual := diagonalDifference(tt.arg)
			if actual != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, actual)
			}
		})
	}
}
