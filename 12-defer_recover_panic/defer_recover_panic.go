package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("This prints last")
	defer fmt.Println("This prints before last")
	fmt.Println("This prints first")
	fmt.Println("This prints second")
	functionWithRecover()
	fmt.Println("This will print because of the recover?")
	fmt.Println("End of main function")
}

func functionWithRecover() {
	defer func() {
		recover()
	}()
	functionWithPanic()
	fmt.Println("This will not print because of the panic")
}

func functionWithPanic() {
	print("About to panic...\n")
	panic("This is a panic message from functionWithPanic. Will not print because of recover")
}

//cd 12-defer
//go run defer.go
