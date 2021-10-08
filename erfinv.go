package math32

import "math"

func Erfinv(x float32) float32 {
	return float32(math.Erfinv(float64(x)))
}
