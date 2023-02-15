package main

import (
	"fmt"
)

func main() {
	var x float64
	var y float64

	x = 90.001
	y = 80.04

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v. type of %T\n", y, y)

	mean := (x + y) / 2
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}
