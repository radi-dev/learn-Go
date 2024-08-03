package main

import "fmt"

func simpleIf(a, b string) string {
	if a == b {
		return "a is b"
	} else {
		return "Elsed"
	}
}

func statementIf(a, b string) string {
	if b := "nom"; a == b {
		return "a is b"
	}
	return "Nothing"
}

func main() {
	fmt.Println(simpleIf("nom", "noma"))
	fmt.Println(statementIf("nom", "noma"))
}

//cd 9-for
//go run if.go
