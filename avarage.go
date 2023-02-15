package main

import (
	"fmt"
)

func main() {
	var x int
	var y int

	x = 90
	y = 80

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v. type of %T\n", y, y)

	mean := (x + y) / 2
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}
