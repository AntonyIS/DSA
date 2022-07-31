package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := binarySearch(nums, 20)
	verify(result)

}

func binarySearch(list []int, target int) int {
	first, last := 0, len(list)-1

	for first <= last {
		midpoint := (first + last) / 2
		if list[midpoint] == target {
			return midpoint
		} else {
			if list[midpoint] > target {
				last = midpoint - 1
			} else {
				first = midpoint + 1
			}
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
