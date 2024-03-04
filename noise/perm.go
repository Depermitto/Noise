package noise

import (
	"math/rand"
	"sync/atomic"
)

const PermTableCap = 256

var perm = atomic.Value{}

// P is a lazy getter for the permutation table.
func P(i int) int {
	if p := perm.Load(); p != nil {
		return p.([]int)[i]
	}

	wrapping := rand.Perm(PermTableCap)
	p := append(wrapping, wrapping...)
	perm.Store(p)
	return p[i]
}
