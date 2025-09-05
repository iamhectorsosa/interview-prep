package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{
			s:    "()",
			want: true,
		},
		{
			s:    "()[]{}",
			want: true,
		},
		{
			s:    "(]",
			want: false,
		},
		{
			s:    "([])",
			want: true,
		},
		{
			s:    "([)]",
			want: false,
		},
		{
			s:    ")(){}",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%q expected to have valid parentheses is %t", tt.s, tt.want), func(t *testing.T) {
			got := isValid(tt.s)
			assert.Equal(t, tt.want, got, "got %t want %t", got, tt.want)
		})
	}
}
