package main

import (
	"fmt"

	"rsc.io/quote/v4"
)

func main() {
    fmt.Println(quote.Opt())
}

//cd 6-external_packages
//go mod init xt_pkgs
//go mod tidy
//go run .
//returns: "If a program is too slow, it must have a loop."