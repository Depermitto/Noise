package worley

import (
	"github.com/Depermitto/noise/noise/chaos"
	"image"
	"math"
	"slices"
)

type Worley struct {
	Cells    map[coords]coords
	cellSize float64
}

func Make(cellSize int, bounds image.Rectangle) Worley {
	cells := make(map[coords]coords)
	for x := 0; x < bounds.Dx(); x += cellSize {
		for y := 0; y < bounds.Dy(); y += cellSize {
			cells[coord(x, y)] = coord(
				chaos.Rand().IntN(cellSize)+x,
				chaos.Rand().IntN(cellSize)+y,
			)
		}
	}
	return Worley{cellSize: float64(cellSize), Cells: cells}
}

func (w Worley) Noise(x float64, y float64) float64 {
	cur := coord(x, y)

	var dist []float64
	for _, i := range []float64{-1, 0, 1} {
		for _, j := range []float64{-1, 0, 1} {
			neighbour := coord(
				math.Floor((x+i*w.cellSize)/w.cellSize)*w.cellSize,
				math.Floor((y+j*w.cellSize)/w.cellSize)*w.cellSize,
			)

			feat, ok := w.Cells[neighbour]
			if ok {
				dist = append(dist, cur.euclidSquared(feat))
			}
		}
	}
	slices.Sort(dist)
	return math.Sqrt(dist[0])
}
