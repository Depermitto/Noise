package white

import "github.com/Depermitto/noise/noise/chaos"

type White struct{}

func (w White) Noise(float64, float64) float64 {
	return chaos.Rand().Float64()
}
