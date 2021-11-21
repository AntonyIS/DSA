package main

import (
	"errors"
	"fmt"
)

func main() {
	sorted_list := []int{1, 2, 4, 5, 6, 7, 8, 9}
	unsorted_list := []int{9, 2, 6, 1, 3, 7, 5}

	res := merge_sort(unsorted_list)
	fmt.Println(res)
	// linear_search
	ok, err := linear_search(sorted_list, 9)
	fmt.Println("LINEAR SEARCH")
	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println(ok)
	}
	// binary_search
	fmt.Println("BINARY SEARCH")
	pass, err := binary_search(sorted_list, 9)

	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println(pass)
	}

	// binary_search
	fmt.Println("RECURSIVE BINARY  SEARCH")
	check, err := recursive_binary_search(sorted_list, 9)

	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println(check)
	}
}

func linear_search(list []int, target int) (bool, error) {
	/*
		Searches for target value in a list
		returns true if found else error
		Takes 0(n) time
	*/

	// Loop through list
	for _, i := range list {
		if i == target {
			return true, nil
		}
	}
	// return error , target value not found
	return false, errors.New("target value not found")

}

func binary_search(list []int, target int) (bool, error) {
	/*
		Searches for target value in a sorted list by dividing list into sublist
		Discards the sublist that does not have the target value
		return true if found, else err
	*/
	first := 0
	last := len(list) - 1

	// Use while loop to go through list
	for first <= last {

		// get the midpoint if the list
		midpoint := (first + last) / 2

		// Check if value at midpoint is equal to the target value
		if list[midpoint] == target {
			return true, nil
		} else {
			// compare value at midpoint with target value
			if list[midpoint] >= target {
				// Discard the right side of the list and re-assign last value
				last = midpoint - 1
			} else {
				// Discard the left side of the list and re-assign first value
				first = midpoint + 1
			}
		}
	}
	return false, errors.New("target value not found")
}

func recursive_binary_search(list []int, target int) (bool, error) {
	/*
		Searches for target value in a sorted list by dividing list into sublist
		recursive_binary_search calls itself if target value is not found
		Discards the sublist that does not have the target value
		return true if found, else err
	*/
	if len(list) == 0 {
		return false, errors.New("target value not found")
	} else {
		midpoint := len(list) / 2
		if list[midpoint] == target {
			return true, nil
		} else {

			if list[midpoint] > target {
				// Discard the right side of the list
				return recursive_binary_search(list[:midpoint], target)
			} else {
				// Discard the right side of the list
				return recursive_binary_search(list[midpoint+1:], target)
			}

		}

	}

}

func merge_sort(list []int) []int {
	/*
		divide: divid a list into sublist
		conquer : sort the sublist in the previous step
		combine : merges the sublist
	*/

	if len(list) <= 1 {
		return list
	} else {
		left_half, right_half := split(list)
		left := merge_sort(left_half)
		right := merge_sort(right_half)
		return merge(left, right)
	}
}

func split(list []int) ([]int, []int) {
	/*
	   Divide unsorted list at midpoint into sublist
	   returns two sublist - left and right
	*/
	midpoint := len(list) / 2
	left := list[:midpoint]
	right := list[midpoint:]

	return left, right
}

func merge(left []int, right []int) []int {
	var merged_list []int
	j := 0
	i := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged_list = append(merged_list, left[i])
			i += 1
		} else {
			merged_list = append(merged_list, right[j])
			j += 1
		}
	}

	for i < len(left) {
		merged_list = append(merged_list, left[i])
		i += 1
	}
	for j < len(right) {
		merged_list = append(merged_list, right[j])
		j += 1
	}
	return merged_list

}
