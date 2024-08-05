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
}

/*
- Cannot be resized
- An array's length is part of its type
- The type [n]T is an array of n values of type T
*/

//cd 15-arrays
//go run arrays.go