package fbm

import "github.com/Depermitto/noise/noise"

type Modulator interface {
	Modulate(float64, float64, noise.Maker) float64
}
