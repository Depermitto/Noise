package noise

// Maker is standard interface for all structs being able to produce noise
// based on given (x, y) coordinates.
type Maker interface {
	// Noise should return a value in range (0, 1).
	Noise(x float64, y float64) float64
}
