package main

import "fmt"

type IHuman interface {
	Describe()
	TransfusePointerBlood(string) string
}

type Human struct {
	height      float64
	weight      float64
	blood_group string
}

func (person Human) Describe() {
	fmt.Printf("I am a human of height %v and weight %v.\nMy blood group is %v.\n", person.height, person.weight, person.blood_group)
}

func (person *Human) TransfusePointerBlood(target string) string {
	switch target {
	case person.blood_group:
		fmt.Println("Nothing to change here!")
	default:
		fmt.Println("Old blood group:", person.blood_group)
		person.blood_group = target
		fmt.Println("New blood group:", person.blood_group)

	}
	return person.blood_group
}

func main() {
	var m IHuman
	ovia := &Human{34, 56, "o-"}

	m = ovia //No failure shows they are of the same type

	fmt.Printf("%T,%v\n", m, ovia.TransfusePointerBlood("m"))

}

//cd 23-Interfaces
//go run interfaces.go
