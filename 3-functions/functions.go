package main

import (
	"fmt"
	"math"
)

// Basic
func difference(a int, b int) float64 {
	return math.Abs(float64(a - b))
}

// multiple return
func subtractAndSum(a, b int) (int, int) {
	return a - b, a + b
}

// named return
func eval(a, b int) (subtract, sum int) {
	subtract = a - b
	sum = a + b
	return
}

func main() {
	fmt.Println(difference(7, 40))
	fmt.Println(subtractAndSum(7, 40))
	fmt.Println(eval(7, 40))
}

//cd 3-functions
//go run functions.go
