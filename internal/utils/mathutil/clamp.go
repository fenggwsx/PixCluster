package mathutil

import "math"

func Clamp(value, lowerBound, upperBound float64) float64 {
	return math.Max(lowerBound, math.Min(upperBound, value))
}
