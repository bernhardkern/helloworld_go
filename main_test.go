package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	result := square(2)

	assert.EqualValues(t, 4, result)
}

func BenchmarkSquare(b *testing.B) {
	
}
