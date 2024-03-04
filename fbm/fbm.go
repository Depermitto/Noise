package fbm

import "github.com/Depermitto/noise/noise"

type ConfigFunc func(fbm *Fbm)
type Fbm struct {
	freq,
	ampl,
	persistence,
	lacunarity float64
	octaves int
}

func (f *Fbm) Modulate(x float64, y float64, maker noise.Maker) float64 {
	res := 0.0
	ampl := f.ampl
	freq := f.freq

	for octave := 0; octave < f.octaves; octave++ {
		res += ampl * maker.Noise(freq*x, freq*y)

		ampl *= f.persistence
		freq *= f.lacunarity
	}
	return res
}

// New constructs a new *Fbm, setting parameters to reasonable default values:
// freq:        0.005,
// ampl:        1,
// octaves:     1,
// persistence: 0.5,
// lacunarity:  2.0.
// Accepts multiple ConfigFunc to customize instead of default values.
func New(f ...ConfigFunc) *Fbm {
	fbm := &Fbm{
		freq:        0.005,
		ampl:        1,
		octaves:     1,
		persistence: 0.5,
		lacunarity:  2.0,
	}
	for _, f := range f {
		f(fbm)
	}
	return fbm
}

// WithOctaves sets amount of octaves for fbm, this effectively means
// amount of times each noise will run on top of each other. Must be
// greater or equal 1.
func WithOctaves(octaves int) ConfigFunc {
	return func(fbm *Fbm) {
		fbm.octaves = octaves
	}
}

// WithFreq sets frequency of fbm, should be a sufficiently low number.
func WithFreq(freq float64) ConfigFunc {
	return func(fbm *Fbm) {
		fbm.freq = freq
	}
}

// WithAmpl sets amplitude of fbm.
func WithAmpl(ampl float64) ConfigFunc {
	return func(fbm *Fbm) {
		fbm.ampl = ampl
	}
}

// WithPersistence sets persistence of fbm, must be value in range (0, 1).
// Persistence is strength of each cascading amplitude per octave.
func WithPersistence(persistence float64) ConfigFunc {
	return func(fbm *Fbm) {
		fbm.persistence = persistence
	}
}

// WithLacunarity sets lacunarity of fbm, must be value greater than 1.
// Lacunarity is strength of each cascading frequency per octave.
func WithLacunarity(lacunarity float64) ConfigFunc {
	return func(fbm *Fbm) {
		fbm.lacunarity = lacunarity
	}
}
