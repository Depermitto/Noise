package noise

import "math/rand"

type White struct{}

func (w White) Noise(float64, float64) float64 {
	return rand.Float64()
}
