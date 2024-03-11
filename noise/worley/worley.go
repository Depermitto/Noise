package worley

import (
	"github.com/Depermitto/noise/noise/chaos"
	"math"
	"slices"
)

type Worley struct {
	Cells map[coords]coords
	tes   bool
}

func Make(tes bool) Worley {
	return Worley{
		Cells: make(map[coords]coords),
		tes:   tes,
	}
}

func (w Worley) Noise(x float64, y float64) float64 {
	cur := coord(x, y)
	x, y = math.Floor(x), math.Floor(y)

	dist := make([]float64, 0, 9)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			neighbour := coord(
				x+float64(i),
				y+float64(j),
			)

			if neighbour.x < 0 || neighbour.y < 0 {
				continue
			}

			feat, ok := w.Cells[neighbour]
			if !ok {
				w.Cells[neighbour] = coord(
					chaos.Rand().Float64()+neighbour.x,
					chaos.Rand().Float64()+neighbour.y,
				)
			}
			dist = append(dist, cur.euclidSquared(feat))
		}
	}

	if w.tes {
		slices.Sort(dist)
		return min(math.Sqrt(dist[1])-math.Sqrt(dist[0]), 1)
	}

	minDist := slices.Min(dist)
	return min(math.Sqrt(minDist), 1)
}
