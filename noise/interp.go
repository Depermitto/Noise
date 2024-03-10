package noise

type Interp int

const (
	Linear Interp = iota
	Cubic
)

func (i Interp) Interpolate(t, a, b float64) float64 {
	switch i {
	case Linear:
		return a + t*(b-a)
	case Cubic:
		return a + (b-a)*(3.0-t*2.0)*t*t
	default:
		return 0.0
	}
}
