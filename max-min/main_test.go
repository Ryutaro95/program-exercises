package main

import (
	"testing"
)

var cases = []struct {
	desc     string
	k        int32
	arr      []int32
	expected int32
}{
	{desc: "Test 1", k: 2, arr: []int32{1, 4, 7, 2}, expected: 1},
	{desc: "Test 2", k: 4, arr: []int32{1, 2, 3, 4, 10, 20, 30, 40, 100, 200}, expected: 3},
	{desc: "Test 3", k: 3, arr: []int32{10, 100, 300, 200, 1000, 20, 30}, expected: 20},
	{desc: "Test 4", k: 4, arr: []int32{1, 2, 3, 4, 10, 20, 30, 40, 100, 200}, expected: 3},
	{desc: "Test 5", k: 2, arr: []int32{1, 2, 1, 2, 1}, expected: 0},
}

func TestMaxMinV2(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			actual := maxMinV2(tt.k, tt.arr)
			if actual != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, actual)
			}
		})
	}
}

func TestMaxMin(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			actual := maxMin(tt.k, tt.arr)
			if actual != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, actual)
			}
		})
	}
}

func TestMaxMinReversion(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			actual := maxMinReverion(tt.k, tt.arr)
			if actual != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, actual)
			}
		})
	}
}
