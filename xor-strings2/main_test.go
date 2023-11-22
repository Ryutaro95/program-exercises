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
			actual, err := stringsXOR(tC.s, tC.t)
			if err != nil {
				t.Errorf("Expected no error, but got %v", err)
			}
			if actual != tC.expected {
				t.Errorf("expected %s, but got %s.", tC.expected, actual)
			}
		})
	}
}

func TestStringXorErrorCase(t *testing.T) {
	testCases := []struct {
		desc, s, t, expected string
	}{
		{
			desc:     "expect an error if the number of letters in s and t are different",
			s:        "101",
			t:        "00101",
			expected: "",
		},
		{
			desc:     "expect an error if s or t is an empty string",
			s:        "",
			t:        "",
			expected: "",
		},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			actual, err := stringsXOR(tC.s, tC.t)
			if err == nil {
				t.Error("expected an error, but got nil")
			}
			if actual != tC.expected {
				t.Errorf("expected empty string, but got %s.", actual)
			}
		})
	}
}
