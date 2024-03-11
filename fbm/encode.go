package fbm

import (
	"github.com/Depermitto/noise/noise"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func (f *Fbm) EncodeImage(maker noise.Maker, imageBounds image.Rectangle, filename string) {
	img := image.NewGray(imageBounds)
	bounds := img.Bounds()
	for x := range bounds.Dx() {
		for y := range bounds.Dy() {
			noi := f.Modulate(maker, float64(x), float64(y))
			intensity := uint8(noi * math.MaxUint8)
			img.SetGray(x, y, color.Gray{Y: intensity})
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return
	}

	err = png.Encode(file, img)
	if err != nil {
		log.Fatalln(err)
	}
}
