package solutions

import "errors"

func Binary_search(list []int, target int) (bool, error) {
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
