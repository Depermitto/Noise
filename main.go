package main

import (
	"github.com/Depermitto/noise/fbm"
	"github.com/Depermitto/noise/noise"
	"github.com/Depermitto/noise/noise/perlin"
	"image"
)

const side = 2 << 8

var bounds = image.Rect(0, 0, side, side)

func main() {
	// chaos.SetSeed(6)
	mod := fbm.New(
		fbm.WithFreq(0.005),
		fbm.WithOctaves(7),
		fbm.WithAmpl(0.5),
	)
	// wor := worley.Make(true)

	mod.EncodeImage(perlin.Make(noise.Cubic), bounds, "worley.png")
}

