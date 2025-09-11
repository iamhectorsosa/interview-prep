package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{},
			target: 0,
			want:   0,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("for %v with a target of %d, the insert position should be %d", tt.nums, tt.target, tt.want), func(t *testing.T) {
			got := searchInsert(tt.nums, tt.target)
			assert.Equal(t, tt.want, got, "got %d want %d", got, tt.want)
		})
	}
}
