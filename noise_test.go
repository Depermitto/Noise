package main

import (
	"github.com/Depermitto/noise/fbm"
	"github.com/Depermitto/noise/noise"
	"testing"
)

func BenchmarkGenerateAndEncode(b *testing.B) {
	gen := Generator{
		maker: noise.Perlin{},
		fbm:   fbm.New(),
	}

	for i := 0; i < b.N; i++ {
		encodeImage(gen, "perlin.png")
	}
}
