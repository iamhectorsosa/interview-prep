package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want bool
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("given root %v", tt.root), func(t *testing.T) {
			got := isSymmetric(tt.root)
			assert.Equal(t, tt.want, got, "got %t want %t", got, tt.want)
		})
	}
}
