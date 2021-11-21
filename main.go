package main

import (
	"errors"
	"fmt"
)

func main() {
	mylist := []int{1, 2, 4, 5, 6, 7, 8, 9}
	// linear_search
	ok, err := linear_search(mylist, 9)
	fmt.Println("LINEAR SEARCH")
	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println(ok)
	}
	// binary_search
	fmt.Println("BINARY SEARCH")
	pass, err := binary_search(mylist, 9)

	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println(pass)
	}

	// binary_search
	fmt.Println("RECURSIVE BINARY  SEARCH")
	check, err := recursive_binary_search(mylist, 9)

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
