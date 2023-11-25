package main

import "fmt"

func strStr(haystack string, needle string) int {
	// Code here
	if len(haystack) == 0 && len(needle) == 0 {
		return 0
	} else if len(haystack) < len(needle) {
		return -1
	}

	n := len(needle)
	for i := range haystack {
		if haystack[i:n+i] == needle {
			return i
		}
		if n+i >= len(haystack) {
			break
		}
	}

	return -1
}

func main() {
	// Code here
	haystack := "leetcode"
	needle := "leeto"
	result := strStr(haystack, needle)
	fmt.Println(result)

	haystack2 := "sadbutsad"
	needle2 := "sad"
	result2 := strStr(haystack2, needle2)
	fmt.Println(result2)
}
