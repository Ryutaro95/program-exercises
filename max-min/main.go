package main

import (
	"math"
	"sort"
)

// Constraints
// 2 ≦ n ≦ 10^5
// 2 ≦ k ≦ n
// 0 ≦ arr[i] ≦ 10^9

func maxMinReverion(k int32, arr []int32) int32 {
	minUnfairness := int32(math.MaxInt32)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	// for start := 0; end := start + int(k) - 1; end < len(arr); start++ {
	for i := 0; i < len(arr)-(int(k)-1); i++ {
		minValue := arr[i]
		maxValue := arr[int(k)+i-1]
		difference := maxValue - minValue
		if minUnfairness > difference {
			minUnfairness = difference
		}
	}

	return minUnfairness
}

func maxMin(k int32, arr []int32) int32 {
	// arr の値は一意の値とは限らない
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	minUnfairness := math.MaxInt
	for i := range arr {
		if len(arr)-1 < int(k)+i {
			break
		}
		window := arr[i : int(k)+i]
		minValue, maxValue := findMinMax(window)
		if minUnfairness > int(maxValue-minValue) {
			minUnfairness = int(maxValue - minValue)
		}
	}
	return int32(minUnfairness) // 不公平 = max(4, 7) - min(4, 7) = 4, 7 = 3
}

func findMinMax(arr []int32) (int32, int32) {
	minValue := math.MaxInt
	maxValue := math.MinInt
	for _, val := range arr {
		if int(val) < minValue {
			minValue = int(val)
		}
		if int(val) > maxValue {
			maxValue = int(val)
		}
	}
	return int32(minValue), int32(maxValue)
}

func main() {
	k := int32(4)
	arr := []int32{1, 2, 3, 4, 10, 20, 30, 40, 100, 200}
	maxMinReverion(k, arr)
}
