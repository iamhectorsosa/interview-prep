package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{
			haystack: "",
			needle:   "",
			want:     0,
		},
		{
			haystack: "sadbutsad",
			needle:   "sad",
			want:     0,
		},
		{
			haystack: "leetcode",
			needle:   "leeto",
			want:     -1,
		},
		{
			haystack: "hello",
			needle:   "ll",
			want:     2,
		},
		{
			haystack: "mississippi",
			needle:   "issi",
			want:     1,
		},
		{
			haystack: "hector",
			needle:   "or",
			want:     4,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("if %q in %q, the index of the first occurance is %d", tt.needle, tt.haystack, tt.want), func(t *testing.T) {
			got := strStr(tt.haystack, tt.needle)
			assert.Equal(t, tt.want, got, "got %d want %d", got, tt.want)
		})
	}
}
