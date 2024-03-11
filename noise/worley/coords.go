package worley

type coords struct {
	x, y float64
}

func coord[T ~int | ~float64](x, y T) coords {
	return coords{float64(x), float64(y)}
}

func (c coords) euclidSquared(other coords) float64 {
	x := c.x - other.x
	y := c.y - other.y
	return x*x + y*y
}
