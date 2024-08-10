package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	var wordCountMap map[string]int = make(map[string]int)
	for _, word := range words {
		wordCountMap[word] += 1
	}
	return wordCountMap
}

func main() {
	wc.Test(WordCount)
}

//cd 20-ex-3-maps
//go mod init wordCount
//go mod tidy
//go run .
