package noise

import (
	"math/rand"
	"sync/atomic"
)

const PermTableCap = 256

var perm = atomic.Value{}

func P(i int) int {
	if perm.Load() == nil {
		wrapping := rand.Perm(PermTableCap)
		perm.Store(append(wrapping, wrapping...))
	}
	return perm.Load().([]int)[i]
}

type Maker interface {
	Noise(x float64, y float64) float64
}
