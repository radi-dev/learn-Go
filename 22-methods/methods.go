package main

import "fmt"

type Human struct {
	height      float64
	weight      float64
	blood_group string
}

func (person Human) Describe() {
	fmt.Printf("I am a human of height %v and weight %v.\nMy blood group is %v.\n", person.height, person.weight, person.blood_group)
}

// Method with value reciever: Cannot modify the receiver, can take either value or pointer as args
func (person Human) TransfuseBlood(target string) string {
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

// Method with pointer reciever: Can  modify the receiver, can take either value or pointer as args
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

// Function with value reciever: Cannot modify the receiver, can only accept pointer arg.. and vice versa
func TransfuseBloodFunc(person Human, target string) {
	switch target {
	case person.blood_group:
		fmt.Println("Nothing to change here!")
	default:
		fmt.Println("Old blood group:", person.blood_group)
		person.blood_group = target
		fmt.Println("New blood group:", person.blood_group)

	}
}

// Function with pointer reciever: Can modify the receiver, can only accept pointer arg.. and vice versa
func TransfusePointerBloodFunc(person *Human, target string) {
	switch target {
	case person.blood_group:
		fmt.Println("Nothing to change here!")
	default:
		fmt.Println("Old blood group:", person.blood_group)
		person.blood_group = target
		fmt.Println("New blood group:", person.blood_group)

	}
}

func main() {
	cheta := Human{34, 56, "o-"}

	fmt.Print("\n")
	cheta.Describe()
	cheta.TransfuseBlood("B+")
	cheta.Describe() //Shows that receiver remains unmodified

	fmt.Print("-------\n")
	cheta.TransfusePointerBlood("B-")
	cheta.Describe() //Shows that receiver gets modified

	fmt.Print("-------\n")
	TransfuseBloodFunc(cheta, "A-")
	cheta.Describe() //Shows that receiver remains unmodified

	fmt.Print("-------\n")
	TransfusePointerBloodFunc(&cheta, "A-")
	cheta.Describe() //Shows that receiver gets modified

}

//cd 22-methods
//go run methods.go
