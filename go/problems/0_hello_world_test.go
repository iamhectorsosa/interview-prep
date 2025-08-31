package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	assert.Equal(t, want, got, "got %q want %q", got, want)
}
