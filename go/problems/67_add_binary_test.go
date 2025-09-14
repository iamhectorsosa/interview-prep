package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want string
	}{
		{
			a:    "11",
			b:    "1",
			want: "100",
		},
		{
			a:    "11",
			b:    "11",
			want: "110",
		},
		{
			a:    "1010",
			b:    "1011",
			want: "10101",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("given %q and %q binary sum we want %q", tt.a, tt.b, tt.want), func(t *testing.T) {
			got := addBinary(tt.a, tt.b)
			assert.Equal(t, tt.want, got, "got %q want %q", got, tt.want)
		})
	}
}
