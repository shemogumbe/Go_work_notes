package main

import (
	"fmt"
)

func main() {
	var x int
	var y int

	x = 51
	y = 41

	if x > y && x < 51 {
		fmt.Printf("X is greater  but within range")

	} else {
		fmt.Printf("Y is greater than X")

	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		// read as fraction is a/b and frac is greater than 0.5
		fmt.Println("A is more than half of b")
	}
}
