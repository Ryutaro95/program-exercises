package main

import "fmt"

func removeDuplicates(nums []int) int {
	// return the length of the array after removing duplicates
	if len(nums) == 0 {
		return 0
	}
	var count int
	for i := 1; i < len(nums); i++ {
		if nums[count] != nums[i] {
			nums[count+1] = nums[i]
			count++
		}
	}
	fmt.Println(nums)
	return count + 1
}

func main() {
}
