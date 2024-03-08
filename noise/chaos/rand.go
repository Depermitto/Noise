package chaos

import (
	"math/rand/v2"
)

var globalRand *rand.Rand

func Rand() *rand.Rand {
	if globalRand == nil {
		SetSeed(rand.Uint64())
	}
	return globalRand
}

func SetSeed(seed uint64) {
	hi := seed >> 32
	lo := seed << 32
	globalRand = rand.New(rand.NewPCG(hi, lo))
}
