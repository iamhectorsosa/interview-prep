package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		want string
	}{
		{
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			strs: []string{},
			want: "",
		},
		{
			strs: []string{"cir", "car"},
			want: "c",
		},
		{
			strs: []string{"baab", "bacb", "b", "cbc"},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("for the array of %v we expect the prefix %q", tt.strs, tt.want), func(t *testing.T) {
			got := longestCommonPrefix(tt.strs)
			assert.Equal(t, tt.want, got, "got %q want %q", got, tt.want)
		})
	}
}
