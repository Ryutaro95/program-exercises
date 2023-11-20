package main

import (
	"reflect"
	"testing"
)

func TestMatchingStringsHashMapVer(t *testing.T) {
	cases := []struct {
		description string
		strs        []string
		queries     []string
		expected    []int32
	}{
		{
			"expect to be able to count how may times the queries slice string appears in he strs slice string",
			[]string{"aba", "baba", "aba", "xzxb"},
			[]string{"aba", "xzxb", "ab"},
			[]int32{2, 1, 0},
		},
		{
			"expect to be able to count duplicate strings of queris",
			[]string{"aba", "baba", "aba", "xzxb"},
			[]string{"aba", "xzxb", "ab", "aba"},
			[]int32{2, 1, 0, 2},
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			result := matchingStringsHashMapVer(tt.strs, tt.queries)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %d, but got %d", tt.expected, result)
			}
		})
	}
}
