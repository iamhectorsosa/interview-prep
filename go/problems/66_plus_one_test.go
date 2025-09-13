package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		digits []int
		want   []int
	}{
		{
			digits: []int{0},
			want:   []int{1},
		},
		{
			digits: []int{1, 2, 3},
			want:   []int{1, 2, 4},
		},
		{
			digits: []int{4, 3, 2, 1},
			want:   []int{4, 3, 2, 2},
		},
		{
			digits: []int{9},
			want:   []int{1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("given %v we want %v", tt.digits, tt.want), func(t *testing.T) {
			got := plusOne(tt.digits)
			assert.Equal(t, tt.want, got, "got %v want %v", got, tt.want)
		})
	}
}
