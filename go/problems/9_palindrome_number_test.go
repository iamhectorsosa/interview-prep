package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		x    int
		want bool
	}{
		{
			x:    121,
			want: true,
		},
		{
			x:    -121,
			want: false,
		},
		{
			x:    10,
			want: false,
		},
		{
			x:    0,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("is palindrome for %d should return %t", tt.x, tt.want), func(t *testing.T) {
			got := isPalindrome(tt.x)
			assert.Equal(t, tt.want, got, "got %t want %t", got, tt.want)
		})
	}
}

func TestIsPalindromeStr(t *testing.T) {
	tests := []struct {
		x    int
		want bool
	}{
		{
			x:    121,
			want: true,
		},
		{
			x:    -121,
			want: false,
		},
		{
			x:    10,
			want: false,
		},
		{
			x:    0,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("is palindrome for %d should return %t", tt.x, tt.want), func(t *testing.T) {
			got := isPalindromeStr(tt.x)
			assert.Equal(t, tt.want, got, "got %t want %t", got, tt.want)
		})
	}
}
