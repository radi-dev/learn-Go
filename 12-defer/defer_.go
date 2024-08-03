package main

import "fmt"

func main() {
	defer fmt.Println("This prints last")
	defer fmt.Println("This prints before last")
	fmt.Println("This prints first")
	fmt.Println("This prints second")
}

//cd 12-defer
//go run defer_.go