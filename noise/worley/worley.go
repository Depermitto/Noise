package worley

import (
	"github.com/Depermitto/noise/noise/chaos"
	"image"
	"math"
	"slices"
)

type Worley struct {
	SeedPts [][2]float64
}

func Make(size int, bounds image.Rectangle) (w Worley) {
	w.SeedPts = make([][2]float64, size)
	for i := range size {
		w.SeedPts[i] = [2]float64{
			float64(chaos.Rand().IntN(bounds.Dx())),
			float64(chaos.Rand().IntN(bounds.Dy())),
		}
	}
	return w
}

func (w Worley) Noise(x float64, y float64) float64 {
	dist := make([]float64, len(w.SeedPts))
	for i, pt := range w.SeedPts {
		dist[i] = math.Sqrt(math.Pow(x-pt[0], 2) + math.Pow(y-pt[1], 2))
	}
	slices.Sort(dist)
	return dist[1] - dist[0]
}
