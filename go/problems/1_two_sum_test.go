package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{},
			target: 0,
			want:   []int{},
		},
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("from %v with target of %d we want %v", tt.nums, tt.target, tt.want), func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			assert.Equal(t, tt.want, got, "got %q want %q", got, tt.want)
		})
	}
}
