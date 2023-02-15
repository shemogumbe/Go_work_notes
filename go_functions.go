package main

import (
	"fmt"
	"math"
)

func main() {
	val := add(3, 4)
	fmt.Println(val)
	fmt.Println("-----------------")
	div, mod := divmod(13, 5)
	fmt.Println("div, Mod", div, mod)
	fmt.Println("----------------------")
	num := 10
	double(num)
	fmt.Println("Doubled number -", num)
	fmt.Println("----------------------")
	doubleAt(&num)
	fmt.Println("Doubled number, by address ", num)
	fmt.Println("-------------------------")
	s1, err := sqrt(-5)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		fmt.Println("The square root", s1)
	}
}

func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func double(n int) int {
	// Does not work,
	n *= 2
	return n
}

func doubleAt(n *int) int {
	// Does not work,
	*n *= 2
	return *n
}

// Returning errors
func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("Square root of negative (%f)", n)
	}
	return math.Sqrt(n), nil
}
