package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVecAdd(t *testing.T) {
	a := Vec(1.0, 1.0)
	b := Vec(2.0, 2.0)

	assert.Equal(t, Vec(3.0, 3.0), a.Add(b))
}

func TestVecAddMix(t *testing.T) {
	a := Vec(1.0, 1.0)
	b := Vec(2.0, 2.0, 2.0)

	assert.Equal(t, Vec(3.0, 3.0), a.Add(b))
	assert.Equal(t, Vec(3.0, 3.0), b.Add(a))
}

func TestVecDimension(t *testing.T) {
	assert.Equal(t, Vec(1.0, 2.0, 3.0, 4.0).Dimension(), 4)
	assert.Equal(t, Vec(1.0, 2.0, 3.0).Dimension(), 3)
	assert.Equal(t, Vec(1.0, 2.0, 3.0, 4.0, 5.0).Dimension(), 5)
}

func TestVecCoordinate(t *testing.T) {
	a := Vec(1.0, 2.0)
	assert.Equal(t, a.Component(0), 1.0)
	assert.Equal(t, a.Component(1), 2.0)
	assert.Equal(t, a.Component(2), 0.0)
}

func TestVecXYZ(t *testing.T) {
	a := Vec(1.0, 2.0, 3.0)
	assert.Equal(t, a.U(), 1.0)
	assert.Equal(t, a.V(), 2.0)
	assert.Equal(t, a.W(), 3.0)
}

func TestVecLen(t *testing.T) {
	assert.Equal(t, Vec(1.0, 2.0, 2.0).Len(), 3.0)
	assert.Equal(t, Vec(2.0, 3.0, 6.0).Len(), 7.0)
	assert.Equal(t, Vec(4.0, 4.0, 7.0).Len(), 9.0)
	assert.Equal(t, Vec(1.0, 4.0, 8.0).Len(), 9.0)
}

func TestVecNorm(t *testing.T) {
	assert.Equal(t, Vec(1.0).Norm(), Vec(1.0))
	assert.Equal(t, Vec(3.0, 4.0).Norm(), Vec(0.6, 0.8))
}

func TestVecDot(t *testing.T) {
	assert.Equal(t, Vec(1.0, 1.0).Dot(Vec(2.0, 3.0)), 5.0)
	assert.Equal(t, Vec(1.0, 1.0).Dot(Vec(-1.0, 1.0)), 0.0)
}
