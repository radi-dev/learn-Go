package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	fmt.Printf("The number is: %v\n----------\n", x)
	z = 1.0
	count := 1
	for {
		_z := z
		z -= (z*z - x) / (2 * z)
		z = math.Round(z*100) / 100

		fmt.Printf("%v. checking: %v\n", count, z)
		if _z == z {
			fmt.Printf("===========\nThe answer is: %v\n===========\n", z)
			break
		}
		if count == 100 {
			break
		}

		count++
	}
	return
}

func main() {
	Sqrt(144)
}

//cd 10-ex-1_sqrt_finder
//go run sqrt.go
