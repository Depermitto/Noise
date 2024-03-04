package noise

type interp int

const (
	Linear interp = iota
	Cubic
)

func (i interp) Interpolate(t, a, b float64) float64 {
	switch i {
	case Linear:
		return a + t*(b-a)
	case Cubic:
		return a + (b-a)*(3.0-t*2.0)*t*t
	default:
		return 0.0
	}
}
