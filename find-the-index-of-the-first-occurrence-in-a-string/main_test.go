package main

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	cases := []struct {
		desc     string
		haystack string
		needle   string
		expected int
	}{
		{desc: "1", haystack: "leetcode", needle: "leeto", expected: -1},
		{desc: "2", haystack: "sadbutsad", needle: "sad", expected: 0},
		{desc: "3", haystack: "hello", needle: "ll", expected: 2},
		{desc: "4", haystack: "aaaaa", needle: "bba", expected: -1},
		{desc: "5", haystack: "", needle: "", expected: 0},
		{desc: "6", haystack: "", needle: "a", expected: -1},
		{desc: "7", haystack: "a", needle: "", expected: 0},
		{desc: "8", haystack: "a", needle: "a", expected: 0},
		{desc: "9", haystack: "mississippi", needle: "issip", expected: 4},
		{desc: "10", haystack: "mississippi", needle: "issipi", expected: -1},
		{desc: "11", haystack: "mississippi", needle: "pi", expected: 9},
		{desc: "12", haystack: "mississippi", needle: "issipi", expected: -1},
		{desc: "13", haystack: "mississippi", needle: "issip", expected: 4},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := strStr(c.haystack, c.needle)
			if got != c.expected {
				t.Errorf("strStr(%q, %q) == %d, want %d", c.haystack, c.needle, got, c.expected)
			}
		})
	}
}
