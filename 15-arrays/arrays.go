package main

import "fmt"

func main() {
	var arr [2]string
	arr[0] = "Ra"
	arr[1] = "di"
	fmt.Println(arr[0] + arr[1])
	fmt.Println(arr)

	blood_code := [5]int{2, 12, 15, 15, 4}
	fmt.Println(blood_code)

	x := [2]string{"Penn", "Teller"}
	y := [...]string{"Penn", "Teller"}
	z := []string{"Penn", "Teller"}  //slice
	fmt.Println("\nx >>", x, "\ny >>", y, "\nz >>", z)

	//arrays are not mutable in Go.
	//To achieve mutation, use slices or pointers
	l := [3]int{}
	n := l
	l[1] = 5
	n[1] = 7
	fmt.Println("\nl >>", l, "\nn >>", n)


}

/*
- Cannot be resized
- An array's length is part of its type
- The type [n]T is an array of n values of type T
*/

//cd 15-arrays
//go run arrays.go
