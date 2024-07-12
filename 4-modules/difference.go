package main

import (
	"math"
)

func difference(a int, b int) float64 {
	return math.Abs(float64(a - b))
}
