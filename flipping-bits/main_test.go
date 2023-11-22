package main

import "testing"

func TestFlippingBits(t *testing.T) {
	cases := []struct {
		description string
		decimal     int64
		expected    int64
	}{
		{"32ビットの最大値を検証", 0, 4294967295},
		{"32ビットの最小値を検証", 4294967295, 0},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			actual := flippingBits(tt.decimal)
			if actual != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, actual)
			}
		})
	}
}
