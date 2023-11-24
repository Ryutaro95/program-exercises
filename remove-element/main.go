package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}

func removeElementV2(nums []int, val int) int {
	i := 0
	n := len(nums)
	for i < n {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}
	fmt.Println(nums)
	return n
}

func main() {
	nums := []int{2, 2, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
}
