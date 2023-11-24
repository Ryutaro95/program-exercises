package main

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	cases := []struct {
		list1    *ListNode
		list2    *ListNode
		expected *ListNode
		desc     string
	}{
		{
			list1:    &ListNode{1, &ListNode{2, nil}},
			list2:    &ListNode{3, &ListNode{4, nil}},
			expected: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
			desc:     "",
		},
		{
			list1:    &ListNode{0, &ListNode{3, &ListNode{8, nil}}},
			list2:    &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{7, &ListNode{8, &ListNode{9, nil}}}}}},
			expected: &ListNode{0, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{7, &ListNode{8, &ListNode{8, &ListNode{9, nil}}}}}}}}},
			desc:     "",
		},
		{
			list1:    nil,
			list2:    nil,
			expected: nil,
			desc:     "",
		},
		{
			list1:    nil,
			list2:    &ListNode{0, nil},
			expected: &ListNode{0, nil},
			desc:     "",
		},
	}
	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			actual := mergeTwoLists(tt.list1, tt.list2)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, actual)
			}
		})
	}
}

func TestMergeTwoListsV2(t *testing.T) {
	cases := []struct {
		list1    *ListNode
		list2    *ListNode
		expected *ListNode
		desc     string
	}{
		{
			list1:    &ListNode{1, &ListNode{2, nil}},
			list2:    &ListNode{3, &ListNode{4, nil}},
			expected: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
			desc:     "",
		},
		{
			list1:    &ListNode{0, &ListNode{3, &ListNode{8, nil}}},
			list2:    &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{7, &ListNode{8, &ListNode{9, nil}}}}}},
			expected: &ListNode{0, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{7, &ListNode{8, &ListNode{8, &ListNode{9, nil}}}}}}}}},
			desc:     "",
		},
		{
			list1:    nil,
			list2:    nil,
			expected: nil,
			desc:     "",
		},
		{
			list1:    nil,
			list2:    &ListNode{0, nil},
			expected: &ListNode{0, nil},
			desc:     "",
		},
	}
	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			actual := mergeTwoListsV2(tt.list1, tt.list2)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, actual)
			}
		})
	}
}
