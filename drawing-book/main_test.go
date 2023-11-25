package main

import "testing"

func TestPageCount(t *testing.T) {
	testCases := []struct {
		desc       string
		pages      int32
		targetPage int32
		expected   int32
	}{
		{
			desc:       "",
			pages:      6,
			targetPage: 2,
			expected:   1,
		},
		{
			desc:       "",
			pages:      4,
			targetPage: 4,
			expected:   0,
		},
		{
			desc:       "",
			pages:      96993,
			targetPage: 70030,
			expected:   13481,
		},
		{
			desc:       "",
			pages:      83246,
			targetPage: 78132,
			expected:   2557,
		},
		{
			desc:       "",
			pages:      18183,
			targetPage: 18042,
			expected:   70,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := pageCount(tC.pages, tC.targetPage)
			if actual != tC.expected {
				t.Errorf("expected %d, but got %d", tC.expected, actual)
			}
		})
	}
}
