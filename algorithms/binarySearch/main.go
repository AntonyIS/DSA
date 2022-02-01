package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	response := binarySearch(nums, 9)
	fmt.Println(response)
}

func binarySearch(nums []int, target int) string {
	// define first and last index
	first, last := 0, len(nums)-1

	for first <= last {
		// Get the midpoint
		midpoint := (first + last) / 2
		// Check if the target value is the midpoint value
		if nums[midpoint] == target {
			// return value at midmidpoint
			return "target value found"
		}
		// Check if target value is between first and midpoint
		if nums[midpoint] >= target {
			// reassign the last value
			last = midpoint - 1
			// Check if target value is between midpoint and last
		} else if nums[midpoint] <= target {
			// reassign the first value
			first = midpoint + 1
		}
	}
	return "target value not found"
}
