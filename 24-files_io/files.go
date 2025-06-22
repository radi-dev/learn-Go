package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func CreateAndWriteFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	file.WriteString("I am Radi.\n\n======== \n\nFile created with Go\n\n\n")
	for i := range 10 {
		for j := range 80 {
			if j == int(math.Pow(float64(i), 2))+2 {
				file.WriteString("✅")
			} else {
				file.WriteString("x")

			}
		}
		file.WriteString("\n")
	}
	fmt.Println("file created: ", file)
}

func ReadFile(fileName string) {
	stream, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	s1 := string(stream)
	fmt.Println(s1)

}

func main() {
	const fileName = "ra.txt"
	CreateAndWriteFile(fileName)
	ReadFile(fileName)

}

//cd 24-files_io
//go run files.go
