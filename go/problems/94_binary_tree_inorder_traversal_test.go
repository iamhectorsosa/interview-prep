package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			want: []int{1, 3, 2},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 9,
						},
					},
				},
			},
			want: []int{4, 2, 6, 5, 7, 1, 3, 9, 8},
		},
		{
			root: nil,
			want: []int{},
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("inorder traversal %v", tt.root), func(t *testing.T) {
			got := inorderTraversal(tt.root)
			assert.Equal(t, tt.want, got, "got %v want %v", got, tt.want)
		})
	}
}
