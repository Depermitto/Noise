package main

import (
	"github.com/Depermitto/noise/fbm"
	"github.com/Depermitto/noise/noise"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

const side = 2 << 8

func main() {
	var (
		mod    = fbm.New()
		perlin = noise.NewPerlin(noise.Linear)
	)

	encodeImage(Generator{perlin, mod}, "perlin.png")
}

func encodeImage(gen Generator, filename string) {
	img := image.NewGray(image.Rect(0, 0, side, side))
	bounds := img.Bounds()
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			intensity := gen.Pix(x, y)
			img.SetGray(x, y, color.Gray{Y: intensity})
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}

	err = png.Encode(f, img)
	if err != nil {
		log.Fatalln(err)
	}
}

type Generator struct {
	maker noise.Maker
	fbm   *fbm.Fbm
}

func (g Generator) Pix(x int, y int) uint8 {
	res := 0.0
	if g.fbm == nil {
		res = g.maker.Noise(float64(x), float64(y))
	} else {
		res = g.fbm.Modulate(float64(x), float64(y), g.maker)
	}
	return uint8(res * math.MaxUint8)
}
