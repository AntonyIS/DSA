package solutions

import "errors"

func Linear_search(list []int, target int) (bool, error) {
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
