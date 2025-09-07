package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{},
			want: 0,
		},
		{
			nums: []int{1, 1, 2},
			want: 2,
		},
		{
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("for %v after removing duplicates we have %d unique items", tt.nums, tt.want), func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			assert.Equal(t, tt.want, got, "got %d want %d", got, tt.want)
		})
	}
}
