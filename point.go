package main

type P struct {
	c []float64
}

func Point(c ...float64) P {
	return P{c}
}

func (p P) Dimension() int {
	return len(p.c)
}

func (p P) Coordinate(i int) float64 {
	if i >= len(p.c) {
		return 0
	}
	return p.c[i]
}

func (a P) Add(b P) P {
	l := min(a.Dimension(), b.Dimension())
	c := []float64{}
	for i := 0; i < l; i++ {
		c = append(c, a.Coordinate(i)+b.Coordinate(i))
	}

	return P{c}
}

func (p P) X() float64 {
	return p.Coordinate(0)
}
func (p P) Y() float64 {
	return p.Coordinate(1)
}
func (p P) Z() float64 {
	return p.Coordinate(2)
}

func (a P) Vec(b P) V {
	l := max(a.Dimension(), b.Dimension())
	o := []float64{}
	for i := 0; i < l; i++ {
		o = append(o, b.Coordinate(i)-a.Coordinate(i))
	}

	return V{o}
}
