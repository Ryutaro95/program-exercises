package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for currentIdx, n := range nums {
		if prevIdx, ok := m[n]; ok {
			return []int{prevIdx, currentIdx}
		}
		m[target-n] = currentIdx
	}

	return []int{}
}

func main() {
}
