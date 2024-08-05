package main

import "fmt"

func main() {
	blood_code := [5]int{2, 12, 15, 15, 4}

	slice_blood_code := blood_code[1:3] //A reference to the array blood_code
	fmt.Println(slice_blood_code)

	var slice_blood_code_2 []int = blood_code[:5] //Another reference to the array blood_code
	fmt.Println(slice_blood_code_2)

	//slice_blood_code[2] =34   ......Error:Out of range
	slice_blood_code[1] = 34
	fmt.Println("blood_code >>", blood_code, "\nslice_blood_code >>", slice_blood_code, "\nslice_blood_code_2 >>", slice_blood_code_2)

	//a slice literal generates its own anonymous array under the hood
	slice_literal_1 := []string{"Radi", "Christi", "Holy", "Molly"}
	fmt.Println("slice_literal_1", slice_literal_1)

	//slices can reference slices
	ovi := slice_literal_1[1:3]
	ovi = append(ovi, "99", "998", "a new ove", "ove tenebrae") //APPEND:creates new underlying array to accomodate capacity of new appendages
	fmt.Println("ovi", ovi)

	//Default for lower bound is 0. For higher bound is length of slice
	ovi_low_default := ovi[:3]
	fmt.Println("ovi_low_default", ovi_low_default)

	ovi_high_default := ovi[1:]
	fmt.Println("ovi_high_default", ovi_high_default)

	ovi_both_default := ovi[:]
	fmt.Println("ovi_both_default", ovi_both_default)
	//------------------------------------------------------------

	fmt.Printf("ovi length is %v and ovi capacity is %v\n", len(ovi), cap(ovi))

	prettyPrint := func(s string, x []int) {
		fmt.Printf("%s len=%d cap=%d %v\n",
			s, len(x), cap(x), x)
	}

	//the built-in make function can be used to create dynamically sized arrays
	a := make([]int, 5)//5 is both len and cap
	prettyPrint("a", a)

	b := make([]int, 0, 12) //0 is len, 12 is cap
	prettyPrint("b", b)

	c := b[:7] //good. within cap of b
	prettyPrint("c", c)

	// d := a[:70]//error. exceeds cap of a/
	// prettyPrint("d", d)

}

/*
SLICE
- Size is dynamic
- More common than array
- The type []T is a slice with elements of type T.
- Is typically a reference to an array
- Mutates its underlying array
- The zero value of a slice is nil.
*/

//cd 16-slices
//go run slices.go
