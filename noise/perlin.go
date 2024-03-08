package noise

import (
	"math"
)

type Perlin struct {
	interp
}

func NewPerlin(interpolation interp) Perlin {
	return Perlin{interp: interpolation}
}

func (p Perlin) Noise(x float64, y float64) float64 {
	xi, yi := int(math.Floor(x))&255, int(math.Floor(y))&255
	xf, yf := x-float64(xi), y-float64(yi)

	var (
		aa = P(P(xi) + yi)
		ba = P(P(xi+1) + yi)
		ab = P(P(xi) + yi + 1)
		bb = P(P(xi+1) + yi + 1)
	)

	u, v := fade(xf), fade(yf)
	return (1 + p.Interpolate(
		u,
		p.Interpolate(v, grad(aa, xf, yf), grad(ab, xf, yf-1)),
		p.Interpolate(v, grad(ba, xf-1, yf), grad(bb, xf-1, yf-1)),
	)) / 2 // moving the range from (-1, 1) to (0, 1).
}

func grad(hash int, x float64, y float64) float64 {
	res := 0.0
	res += float64(1-2*(hash&1)) * x
	res += float64(1-2*(hash&2)>>1) * y
	return res
}

func fade(t float64) float64 {
	return t * t * t * (t*(t*6-15) + 10)
}
