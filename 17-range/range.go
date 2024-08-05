package main

import "fmt"

func forRangeOfSlice(slc []int) {
	fmt.Printf("\nLoop of slice %v.\n", slc)
	for i, v := range slc {
		fmt.Printf("\nIndex is %v\nValue is %v\n", i, v)
	}

	fmt.Printf("\n\nLoop of slice %v. omit value\n", slc)
	for i := range slc {
		fmt.Printf("\nIndex is %v\n", i,)
	}

	fmt.Printf("\n\nLoop of slice %v. skip index\n", slc)
	for _, v := range slc {
		fmt.Printf("\nValue is %v\n", v)
	}

	fmt.Printf("\n\nLoop of slice %v. Skip value\n", slc)
	for i, _ := range slc {
		fmt.Printf("\nIndex is %v\n", i)
	}
}

func main() {
	forRangeOfSlice([]int{5, 9, 77})
}

/*
Contd. from "8-for_loop" module*/

//cd 17-range
//go run range.go
