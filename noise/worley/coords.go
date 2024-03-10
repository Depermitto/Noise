package worley

import "math"

type coords [2]float64

func coord[T ~int | ~float64](x, y T) coords {
	return coords{float64(x), float64(y)}
}

func (c coords) x() float64 {
	return c[0]
}

func (c coords) y() float64 {
	return c[1]
}

func (c coords) euclidSquared(other coords) float64 {
	return math.Pow(c.x()-other.x(), 2) + math.Pow(c.y()-other.y(), 2)
}
