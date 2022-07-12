package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtilMin(t *testing.T) {
	assert.Equal(t, min(3, 4), 3)
	assert.Equal(t, min(4, 3), 3)
}

func TestUtilMax(t *testing.T) {
	assert.Equal(t, max(3, 4), 4)
	assert.Equal(t, max(4, 3), 4)
}
