package main

import (
	"fmt"
)

func main() {
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Println("loons = %v (type %T)\n", loons, loons)

	// Find length
	fmt.Println("Slice length", len(loons))

	//Slice slices - first item
	fmt.Println("First item :", loons[0])

	// slice to get all except first
	fmt.Println("Removed first :", loons[1:])

	// Slice slices- get al except last
	fmt.Println("Removed Last :", loons[:2])

	fmt.Println("--------------------------")
	// Loop and print
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}
	fmt.Println("-----------------------------")
	//Using rangw with slices
	// one item - the index
	for i := range loons {
		fmt.Println(i)
	}

	// Slice with ranges, values

	for i, name := range loons {
		fmt.Println("%s at %d\n", name, i)
	}
	fmt.Println("----------------------------")
	//Ignore index using _
	loons = append(loons, "Elmer")
	for _, name := range loons {
		fmt.Println(name)
	}
}
