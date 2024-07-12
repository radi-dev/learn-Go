package main

import (
	"fmt"
	"math"
)

func difference(a int, b int) float64 {
	return math.Abs(float64(a - b))
}

func main() {
	fmt.Println(difference(7, 40))
}
