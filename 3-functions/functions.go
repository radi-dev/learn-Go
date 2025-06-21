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

// variadic function : Overloaded function that can take a variable number of arguments
func variadicSum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	ns := []int{}
	for i := 1; i <= 100; i++ {
		ns = append(ns, i)
	}
	fmt.Println("Length of ns:", len(ns))

	fmt.Println(difference(7, 40))
	fmt.Println(subtractAndSum(7, 40))
	fmt.Println(eval(7, 40))
	fmt.Println(variadicSum(ns...)) // using range to generate a slice of numbers from 1 to 9
}

//cd 3-functions
//go run functions.go
