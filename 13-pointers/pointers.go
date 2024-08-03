package main

import "fmt"

func pointX() {
	x := 9
	x_pointr := &x // point to x
	fmt.Println("x >>", x, "\n*x_pointr >>", *x_pointr)

	*x_pointr = 12 // set x through the pointer
	fmt.Println("\nx >>", x, "\n*x_pointr >>", *x_pointr)

	x = 24 // set x through the pointer
	fmt.Println("\nx >>", x, "\n*x_pointr >>", *x_pointr)

	*x_pointr = *x_pointr / 2
	fmt.Println("\nx >>", x, "\n*x_pointr >>", *x_pointr)

	*x_pointr = x / 2
	fmt.Println("\nx >>", x, "\n*x_pointr >>", *x_pointr)
}

func main() {
	pointX()
}
