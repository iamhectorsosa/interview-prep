package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "",
			want: 0,
		},
		{
			s:    "Hello World",
			want: 5,
		},
		{
			s:    "   fly me   to   the moon  ",
			want: 4,
		},
		{
			s:    "luffy is still joyboy",
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("for %q the length of last word is %d", tt.s, tt.want), func(t *testing.T) {
			got := lengthOfLastWord(tt.s)
			assert.Equal(t, tt.want, got, "got %d want %d", got, tt.want)
		})
	}
}
