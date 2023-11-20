package main

import "testing"

func TestLonelyInteger(t *testing.T) {
	cases := []struct {
		description string
		arr         []int32
		expected    int32
	}{
		{"expect to be 4", []int32{1, 2, 3, 4, 3, 2, 1}, 4},
		{"expect to get the correct value even if they are regularly lined up", []int32{0, 0, 1, 2, 1}, 2},
		{"expect to get the correct result regardless of the number of element", []int32{18, 49, 15, 30, 56, 20, 49, 67, 82, 69, 84, 63, 93, 87, 66, 17, 38, 32, 17, 32, 94, 66, 67, 63, 48, 64, 84, 82, 87, 18, 79, 64, 79, 15, 71, 94, 59, 5, 22, 59, 4, 98, 81, 4, 98, 73, 69, 88, 30, 81, 48, 56, 71, 20, 93, 22, 73, 5, 88}, 38},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			result := lonelyinteger(tt.arr)
			if result != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, result)
			}
		})
	}
}
