package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
		want int
	}{
		{
			nums: []int{},
			val:  0,
			want: 0,
		},
		{
			nums: []int{3, 2, 2, 3},
			val:  3,
			want: 2,
		},
		{
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("given %v and val %d, we expect %d", tt.nums, tt.val, tt.want), func(t *testing.T) {
			got := removeElement(tt.nums, tt.val)
			assert.Equal(t, tt.want, got, "got %d want %d", got, tt.want)
		})
	}
}
