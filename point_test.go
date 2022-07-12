package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointAdd(t *testing.T) {
	a := Point(1.0, 1.0)
	b := Point(2.0, 2.0)

	assert.Equal(t, Point(3.0, 3.0), a.Add(b))
}

func TestPointAddMix(t *testing.T) {
	a := Point(1.0, 1.0)
	b := Point(2.0, 2.0, 2.0)

	assert.Equal(t, Point(3.0, 3.0), a.Add(b))
	assert.Equal(t, Point(3.0, 3.0), b.Add(a))
}

func TestPointDimension(t *testing.T) {
	assert.Equal(t, Point(1.0, 2.0, 3.0, 4.0).Dimension(), 4)
	assert.Equal(t, Point(1.0, 2.0, 3.0).Dimension(), 3)
	assert.Equal(t, Point(1.0, 2.0, 3.0, 4.0, 5.0).Dimension(), 5)
}

func TestPointCoordinate(t *testing.T) {
	a := Point(1.0, 2.0)
	assert.Equal(t, a.Coordinate(0), 1.0)
	assert.Equal(t, a.Coordinate(1), 2.0)
	assert.Equal(t, a.Coordinate(2), 0.0)
}

func TestPointXYZ(t *testing.T) {
	a := Point(1.0, 2.0, 3.0)
	assert.Equal(t, a.X(), 1.0)
	assert.Equal(t, a.Y(), 2.0)
	assert.Equal(t, a.Z(), 3.0)
}

func TestPointVec(t *testing.T) {
	assert.Equal(t, Point(1.0, 1.0).Vec(Point(3.0, 3.0)), Vec(2.0, 2.0))
}
