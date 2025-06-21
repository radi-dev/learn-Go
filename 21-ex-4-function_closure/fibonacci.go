package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		next := a
		a, b = b, a+b
		return next
	}
}

func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(i, ": ", f())
	}

	// More Closure examples:

	counter := makeCounter()
	fmt.Println("\n\nMore Closure examples:")
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
}

//cd 21-ex-4-function_closure
//go run fibonacci.go
