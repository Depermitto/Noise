package worley

import (
	"github.com/Depermitto/noise/noise/chaos"
	"math"
	"slices"
)

type Worley struct {
	cells map[coords]coords
	tri   bool
}

func Make(tri bool) Worley {
	return Worley{
		cells: make(map[coords]coords),
		tri:   tri,
	}
}

func (w Worley) Noise(x float64, y float64) float64 {
	cur := coord(x, y)

	var dist []float64
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			neighbour := coord(
				math.Floor(x)+float64(i),
				math.Floor(y)+float64(j),
			)

			feat, ok := w.cells[neighbour]
			if !ok {
				w.cells[neighbour] = coord(
					chaos.Rand().Float64()+neighbour.x(),
					chaos.Rand().Float64()+neighbour.y(),
				)
			}
			dist = append(dist, cur.euclidSquared(feat))
		}
	}

	if w.tri {
		slices.Sort(dist)
		return min(math.Sqrt(dist[1])-math.Sqrt(dist[0]), 1)
	}

	minDist := slices.Min(dist)
	return min(math.Sqrt(minDist), 1)
}
