package fbm

import (
	"github.com/Depermitto/noise/noise/worley"
	"image"
	"testing"
)

func BenchmarkWorley512(b *testing.B) {
	wor := worley.Make(false)
	mod := New()
	bounds := image.Rect(0, 0, 512, 512)
	for i := 0; i < b.N; i++ {
		mod.EncodeImage(wor, bounds, "")
	}
}

func BenchmarkWorley1024(b *testing.B) {
	wor := worley.Make(false)
	mod := New()
	bounds := image.Rect(0, 0, 1024, 1024)
	for i := 0; i < b.N; i++ {
		mod.EncodeImage(wor, bounds, "")
	}
}
