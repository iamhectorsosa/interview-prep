package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		head []int
		want []int
	}{
		{
			head: []int{1, 1, 2},
			want: []int{1, 2},
		},
		{
			head: []int{1, 1, 2, 3, 3},
			want: []int{1, 2, 3},
		},
		{
			head: []int{1, 1, 1},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("provided %v, we expect %v", tt.head, tt.want), func(t *testing.T) {
			got := toInt(deleteDuplicates(toList(tt.head)))
			assert.Equal(t, tt.want, got, "got %v want %v", got, tt.want)
		})
	}
}
