package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "",
			want: 0,
		},
		{
			s:    "III",
			want: 3,
		},
		{
			s:    "LVIII",
			want: 58,
		},
		{
			s:    "MCMXCIV",
			want: 1994,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("convert roman %q to an integer", tt.s), func(t *testing.T) {
			got := romanToInt(tt.s)
			assert.Equal(t, tt.want, got, "got %d, want %d", got, tt.want)
		})
	}
}
