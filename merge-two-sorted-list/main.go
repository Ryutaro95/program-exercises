package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var result *ListNode
	if list1.Val <= list2.Val {
		result = &ListNode{list1.Val, mergeTwoLists(list1.Next, list2)}
	} else {
		result = &ListNode{list2.Val, mergeTwoLists(list1, list2.Next)}
	}

	return result
}

func mergeTwoListsV2(l1, l2 *ListNode) *ListNode {
	prehead := new(ListNode)
	prev := prehead

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}

	return prehead.Next
}

func main() {
	l1 := &ListNode{1, &ListNode{2, nil}}
	l2 := &ListNode{3, &ListNode{4, nil}}

	printNodeLists(mergeTwoListsV2(l1, l2))
}

func printNodeLists(n *ListNode) {
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}
