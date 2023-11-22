package main

import "testing"

func Test(t *testing.T) {
	type Args struct {
		k int32
		A []int32
		B []int32
	}

	testCases := []struct {
		desc     string
		args     Args
		expected string
	}{
		{
			desc:     "expect return YES",
			args:     Args{k: 10, A: []int32{2, 1, 3}, B: []int32{7, 8, 9}},
			expected: "YES",
		},
		{
			desc: "expect return YES",
			args: Args{
				94,
				[]int32{84, 2, 50, 51, 19, 58, 12, 90, 81, 68, 54, 73, 81, 31, 79, 85, 39, 2},
				[]int32{53, 102, 40, 17, 33, 92, 18, 79, 66, 23, 84, 25, 38, 43, 27, 55, 8, 19},
			},
			expected: "YES",
		},
		{
			desc: "expect return NO",
			args: Args{
				88,
				[]int32{66, 66, 32, 75, 77, 34, 23, 35},
				[]int32{61, 17, 52, 20, 48, 11, 50, 5},
			},
			expected: "NO",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := twoArrays(tc.args.k, tc.args.A, tc.args.B)
			if actual != tc.expected {
				t.Errorf("expected %s, but got %s.", tc.expected, actual)
			}
		})
	}
}
