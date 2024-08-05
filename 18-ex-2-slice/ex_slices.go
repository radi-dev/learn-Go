package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
}

func main() {
	pic.Show(Pic)
}

//cd 18-ex-2-slice
//go mod init slices_ex
//go mod tidy
//go run .