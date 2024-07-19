package main

import (
	"fmt"
	"strings"
	"time"
)

func forLoop(count int) (sum int) {
	for i := 0; i < count; i++ {
		sum += i
		fmt.Printf("\n\nThis is loop %v.\nThe index is %v.\nCurrent sum is %v.\n\n--------------", i+1, i, sum)
	}
	return
}

// Go has only a for loop. But a while loop behaviour can be achieved.
func whileLoop(count int)(sum int) {
	for sum <= count {
		sum += 1
		fmt.Println(sum)
	}

	return
}

//Danger! Infinite while loop
func infiniteWhile(){
	count:=1
	var exclamation = "!"
	for {
		fmt.Printf("Radi is awesome%s\n\n",strings.Repeat(exclamation,count))
		count++
		time.Sleep(1000000000)
	}
}

func main() {

	forLoop(15)
	whileLoop(12)
	infiniteWhile()

}
