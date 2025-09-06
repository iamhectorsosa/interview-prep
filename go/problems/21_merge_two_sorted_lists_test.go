package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1 []int
		list2 []int
		want  []int
	}{
		{
			list1: []int{1, 2, 4},
			list2: []int{1, 3, 4},
			want:  []int{1, 1, 2, 3, 4, 4},
		},
		{
			list1: []int{},
			list2: []int{},
			want:  []int{},
		},
		{
			list1: []int{},
			list2: []int{0},
			want:  []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("provided %v and %v, we expect %v", tt.list1, tt.list2, tt.want), func(t *testing.T) {
			list1 := toList(tt.list1)
			list2 := toList(tt.list2)
			want := toList(tt.want)
			got := mergeTwoLists(list1, list2)
			assert.Equal(t, want, got, "got %v want %v", got, want)
		})
	}
}

func toList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	var node *ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		node = &ListNode{
			Val:  nums[i],
			Next: node,
		}
	}

	return node
}
