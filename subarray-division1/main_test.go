package main

import "testing"

type Args struct {
	s     []int32
	day   int32
	month int32
}

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		args     Args
		expected int32
	}{
		{
			desc: "expect to be able to divide 7.",
			args: Args{
				s:     []int32{2, 2, 2, 1, 3, 2, 2, 3, 3, 1, 4, 1, 3, 2, 2, 1, 2, 2, 4, 2, 2, 3, 5, 3, 4, 3, 2, 1, 4, 5, 4},
				day:   10,
				month: 4,
			},
			expected: 7,
		},
		{
			desc: "expect to be able to divide 2.",
			args: Args{
				s:     []int32{1, 2, 1, 3, 2},
				day:   3,
				month: 2,
			},
			expected: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := birthday(tC.args.s, tC.args.day, tC.args.month)
			if actual != tC.expected {
				t.Errorf("expected %d, but got %d.", tC.expected, actual)
			}
		})
	}
}
