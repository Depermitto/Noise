package main

import (
	"github.com/Depermitto/noise/fbm"
	"github.com/Depermitto/noise/noise"
	"github.com/Depermitto/noise/noise/chaos"
	"github.com/Depermitto/noise/noise/perlin"
	"image"
)

const side = 2 << 9

var bounds = image.Rect(0, 0, side, side)

func main() {
	chaos.SetSeed(6)
	mod := fbm.New(
		fbm.WithFreq(0.003),
		fbm.WithOctaves(3),
		fbm.WithLacunarity(3.0),
	)
	//wor := worley.Make(true)
	per := perlin.Make(noise.Linear)

	//mod.EncodeImage(wor, bounds, "worley.png")
	//fmt.Println(len(wor.Cells))
	mod.EncodeImage(per, bounds, "perlin.png")
}
