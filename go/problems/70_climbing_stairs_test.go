package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{
			n:    2,
			want: 2,
		},
		{
			n:    3,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("for %d stairs there should be %d ways to the stop", tt.n, tt.want), func(t *testing.T) {
			got := climbStairs(tt.n)
			assert.Equal(t, tt.want, got, "got %d want %d", got, tt.want)
		})
	}
}
