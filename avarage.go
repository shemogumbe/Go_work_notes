package main

import (
	"fmt"
)

func main() {
	n := 9

	switch n {
	case 1:
		fmt.Println("Selected - One")
	case 2:
		fmt.Println("Selected - Two")
	case 3:
		fmt.Println("Selected - Three")
	default:
		fmt.Printf("Selection - Out of range")
	}
}
