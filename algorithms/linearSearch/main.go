package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	response := linearSearch(nums, 9)
	fmt.Println(response)

}

func linearSearch(nums []int, target int) string {
	// Loop through values in the slice
	for _, value := range nums {
		if value == target {
			return "target value found"
		}
	}
	return "taegt value not found"
}
