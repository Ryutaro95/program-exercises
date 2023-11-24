package main

import "testing"

func TestRemoveElement(t *testing.T) {
	cases := []struct {
		nums []int
		val  int
		want int
	}{
		{[]int{1, 2, 3, 5, 4}, 2, 0},
		{[]int{2, 2, 2}, 2, 0},
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for _, c := range cases {
		got := removeElement(c.nums, c.val)
		if got != c.want {
			t.Errorf("removeElement(%v, %d) == %d, want %d", c.nums, c.val, got, c.want)
		}
	}
}

func TestRemoveElementV2(t *testing.T) {
	cases := []struct {
		nums []int
		val  int
		want int
	}{
		{[]int{2, 2, 2}, 2, 0},
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for _, c := range cases {
		got := removeElementV2(c.nums, c.val)
		if got != c.want {
			t.Errorf("removeElement(%v, %d) == %d, want %d", c.nums, c.val, got, c.want)
		}
	}
}
