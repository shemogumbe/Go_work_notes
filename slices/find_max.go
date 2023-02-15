package main

import "fmt"

func main() {
	nums := []int{16, 42, 15, 78, 90}

	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	fmt.Println("Max - ", max)
}
