package main

import (
	"fmt"
	"github.com/Depermitto/noise/fbm"
	"github.com/Depermitto/noise/noise/chaos"
	"github.com/Depermitto/noise/noise/worley"
	"image"
)

const side = 2 << 8

var bounds = image.Rect(0, 0, side, side)

func main() {
	chaos.SetSeed(6)
	mod := fbm.New(fbm.WithFreq(0.005))
	wor := worley.Make(false)

	mod.EncodeImage(wor, bounds, "worley.png")

	for k, v := range wor.Cells {
		fmt.Println(k, v)
	}
}
