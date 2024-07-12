package main

import (
	f "fmt"
	"math"
	"math/rand"
)

func main() {
	f.Println("Arc tan 30:",math.Atan(30))
	f.Print("Random integer between 0 and 50: ")
	f.Println((rand.Intn(50)))
}
