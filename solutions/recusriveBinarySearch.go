package solutions

import "errors"

func Recursive_binary_search(list []int, target int) (bool, error) {
	/*
		Searches for target value in a sorted list by dividing list into sublist
		Recursive_binary_search calls itself if target value is not found
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
				return Recursive_binary_search(list[:midpoint], target)
			} else {
				// Discard the right side of the list
				return Recursive_binary_search(list[midpoint+1:], target)
			}

		}

	}

}
