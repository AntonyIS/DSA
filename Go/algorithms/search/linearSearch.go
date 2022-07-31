package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := linearSearch(nums, 10)
	verify(result)
}

func linearSearch(list []int, target int) int {
	for i := range list {
		if list[i] == target {
			return i
		}
	}
	return 0
}

func verify(result int) {
	if result != 0 {
		fmt.Println("Target found in index: ", result)
	} else {
		fmt.Println("Target not found")
	}
}
