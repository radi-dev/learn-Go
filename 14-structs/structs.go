package main

import "fmt"

type Human struct {
	height      float64
	weight      float64
	blood_group string
}

var radi Human

func main() {
	var christi = Human{height: 75, weight: 67, blood_group: "O+"}

	fmt.Println(Human{})
	fmt.Println(Human{height: 34})
	fmt.Println(Human{34, 45, "O+"})
	fmt.Println(Human{height: 34, weight: 45, blood_group: "O+"})
	fmt.Println(Human{height: 34, weight: 45, blood_group: "O+"})
	fmt.Println(radi)
	radi.height = 110
	fmt.Println(radi)
	fmt.Println(christi.weight)

	radi_pointr := &radi
	radi_pointr.weight = 77
	fmt.Println(radi)

}

//cd 14-structs
//go run structs.go
