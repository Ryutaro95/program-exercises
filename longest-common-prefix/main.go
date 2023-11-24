package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 二分探索を利用するため、最短の文字列の長さを求める
	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	// 二分探索を開始
	low, high := 0, minLen
	for low < high {
		mid := (low + high + 1) / 2

		// 二分探索の結果に基づき、接頭辞の存在を確認
		if isCommonPrefix(strs, mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}

	return strs[0][:low]
}

// 長さがlengthの接頭辞が全ての文字列で共通しているか確認するヘルパー関数
func isCommonPrefix(strs []string, length int) bool {
	prefix := strs[0][:length]
	for _, str := range strs {
		if !strings.HasPrefix(str, prefix) {
			return false
		}
	}
	return true
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(strs)
	fmt.Println(result) // Output: "fl"
}
