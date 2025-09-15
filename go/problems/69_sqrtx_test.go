package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{
			x:    4,
			want: 2,
		},
		{
			x:    8,
			want: 2,
		},
		{
			x:    16,
			want: 4,
		},
		{
			x:    18,
			want: 4,
		},
		{
			x:    0,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("the sqrt of %d rounded to the nearest integer should be %d", tt.x, tt.want), func(t *testing.T) {
			got := mySqrt(tt.x)
			assert.Equal(t, tt.want, got, "got %d want %d", got, tt.want)
		})
	}
}
