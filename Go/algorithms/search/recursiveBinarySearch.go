package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := recursiveBinarySearch(nums, 17)
	verify(result)
}

func recursiveBinarySearch(list []int, target int) bool {
	if len(list) == 0 {
		return false
	} else {
		midpoint := len(list) / 2
		if list[midpoint] == target {
			return true
		} else {
			if list[midpoint] > target {
				return recursiveBinarySearch(list[:midpoint], target)
			} else {
				return recursiveBinarySearch(list[midpoint+1:], target)
			}
		}
	}
}

func verify(result bool) {
	fmt.Printf("Target found : %v \n", result)
}
