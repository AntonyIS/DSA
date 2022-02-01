package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	response := recursiveBinarySearch(nums, 90)
	if response {
		// If response is true
		fmt.Println("target value found")
	} else {
		// if response is false
		fmt.Println("target value not found")
	}
}

func recursiveBinarySearch(nums []int, target int) bool {
	// Check if slice is empty
	if len(nums) == 0 {
		return false
	} else {
		// base condition
		midpoint := len(nums) / 2
		if nums[midpoint] == target {
			return true
		} else if (nums[midpoint]) > target {
			// Work with values from 0 index to midpoint
			return recursiveBinarySearch(nums[:midpoint], target)
		} else if (nums[midpoint]) < target {
			// Work with values from midpoint to the end
			return recursiveBinarySearch(nums[midpoint+1:], target)
		}
	}
	// target value does not exist in the slice
	return false
}
