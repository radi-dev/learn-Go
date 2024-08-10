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

func main() {
	f := fibonacci()
	for i := 0; i < 10000; i++ {
		fmt.Println(i, ": ", f())
	}
}

//cd 20-ex-3-maps
