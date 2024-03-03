package fbm

import "github.com/Depermitto/noise/noise"

type ConfigFunc func(fbm *Fbm)
type Fbm struct {
	freq    float64
	ampl    float64
	octaves int
}

func (f *Fbm) Modulate(x float64, y float64, generator noise.Maker) float64 {
	res := 0.0
	ampl := f.ampl
	freq := f.freq

	for octave := 0; octave < f.octaves; octave++ {
		res += ampl * generator.Noise(freq*x, freq*y)

		ampl *= 0.5
		freq *= 2.0
	}
	return res
}

func New(f ...ConfigFunc) *Fbm {
	fbm := &Fbm{freq: 0.05, ampl: 1, octaves: 1}
	for _, f := range f {
		f(fbm)
	}
	return fbm
}

func WithOctaves(octaves int) ConfigFunc {
	return func(fbm *Fbm) {
		fbm.octaves = octaves
	}
}

func WithFreq(freq float64) ConfigFunc {
	return func(fbm *Fbm) {
		fbm.freq = freq
	}
}

func WithAmpl(ampl float64) ConfigFunc {
	return func(fbm *Fbm) {
		fbm.ampl = ampl
	}
}
