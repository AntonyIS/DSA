package solutions

/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
*/

func ContainsDuplicates(nums []int) bool {
	duplicatsMapper := make(map[int]int)

	for i, value := range nums {
		if _, ok := duplicatsMapper[value]; ok {
			return true
		} else {
			duplicatsMapper[value] = i
		}
	}
	return false
}
