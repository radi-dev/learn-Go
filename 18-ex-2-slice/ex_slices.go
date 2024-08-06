package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic_item := make([][]uint8, dy)
	for i := range pic_item {
		pic_item[i] = make([]uint8, dx)
		for j := range pic_item[i] {
			pic_item[i][j] = (uint8(i) + uint8(j)) / 2
		}
	}

	return pic_item
}

func main() {
	pic.Show(Pic)
}

//cd 18-ex-20-slice
//go mod init slices_ex
//go mod tidy
//go run .
