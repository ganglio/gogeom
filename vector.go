package main

import (
	"math"
)

type V struct {
	c []float64
}

func Vec(c ...float64) V {
	return V{c}
}

func (v V) Dimension() int {
	return len(v.c)
}

func (v V) Component(i int) float64 {
	if i >= len(v.c) {
		return 0
	}
	return v.c[i]
}

func (a V) Add(b V) V {
	l := min(a.Dimension(), b.Dimension())
	c := []float64{}
	for i := 0; i < l; i++ {
		c = append(c, a.Component(i)+b.Component(i))
	}

	return V{c}
}

func (v V) U() float64 {
	return v.Component(0)
}
func (v V) V() float64 {
	return v.Component(1)
}
func (v V) W() float64 {
	return v.Component(2)
}

func (v V) Len() float64 {
	var l float64 = 0
	for i := 0; i < v.Dimension(); i++ {
		l += v.Component(i) * v.Component(i)
	}
	return math.Sqrt(l)
}

func (v V) Norm() V {
	l := v.Len()
	if l > 0 {
		n := []float64{}
		for i := 0; i < v.Dimension(); i++ {
			n = append(n, v.Component(i)/l)
		}
		return V{n}
	}
	return V{}
}

func (a V) Dot(b V) float64 {
	var d float64 = 0
	for i := 0; i < a.Dimension(); i++ {
		d += a.Component(i) * b.Component(i)
	}
	return d
}
