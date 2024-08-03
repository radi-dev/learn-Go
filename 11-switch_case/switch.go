package main

import (
	"fmt"
	"time"
)

func saturdaySwitch() {
	fmt.Println("When's Saturday?")
	switch today := time.Now().Weekday(); time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today - 1:
		fmt.Println("Yesterday.")
	case today - 2:
		fmt.Println("Two days ago.")
	default:
		fmt.Println("Too far away.")
	}
}
func weekdaySwitch() {
	fmt.Println("Weekday or weekend?")
	today := time.Now().Weekday()
	switch {
	case today == time.Saturday:
		fmt.Println("Weekend.")
	case today == time.Sunday:
		fmt.Println("Weekend.")
	case today == time.Wednesday:
		fmt.Println("Midweek.")
	default:
		fmt.Println("Weekday.")
	}
}

func main() {
	saturdaySwitch()
	weekdaySwitch()
}

//cd 11-switch_case
//go run switch.go
