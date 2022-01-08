package solutions

/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
*/

func ContainsDuplicates(nums []int) bool {
	duplicatsMapper := make(map[int]int)

	for _, value := range nums {

		if _, ok := duplicatsMapper[value]; ok {
			// The key already exists in the nums slice, increament the number of visited time
			duplicatsMapper[value] += 1
			// if value has been visited more than 2 times
			if duplicatsMapper[value] == 2 {
				return true
			}
		} else {
			// Mark as once
			duplicatsMapper[value] = 1
		}
	}
	return false
}
