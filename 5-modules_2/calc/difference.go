package calc

import (
	"math"
)

func Difference(a, b int) (float64, int) {
	return math.Abs(float64(a - b)), a - b
}
