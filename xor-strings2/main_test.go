package main

import "testing"

func TestStringXOR(t *testing.T) {
	testCases := []struct {
		desc, s, t, expected string
	}{
		{
			s:        "10101",
			t:        "00101",
			expected: "10000",
		},
		{
			s:        "1101111011000001001001010010001001111110001110100100100100001101101010101001110010000101011110111000",
			t:        "1010001101111011110100010101111110101101100000001001110000110110001100011110100100001010001101010011",
			expected: "0111110110111010111101000111110111010011101110101101010100111011100110110111010110001111010011101011",
		},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			actual := stringsXOR(tC.s, tC.t)
			if actual != tC.expected {
				t.Errorf("expected %s, but got %s.", tC.expected, actual)
			}
		})
	}
}
