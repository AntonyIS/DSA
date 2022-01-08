package solutions

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/

func TwoSum(nums []int, target int) []int {
	// Solution 1 : 0(n2)
	// for i := 0; i < len(nums)-1; i++ {
	// 	for j := 0; j < len(nums); j++ {
	// 		if nums[i]+nums[j] == target {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }

	// Solution 2 : 0(1)
	// Map to store complement keys
	mapper := make(map[int]int)
	for i, value := range nums {
		// Get complement
		complement := target - value
		if _, ok := mapper[complement]; ok {
			// Return the index of the 1st and 2nd value that sum up to target
			return []int{mapper[value], i}
		}
		// Store value as the key and index as a value for reference
		mapper[value] = i

	}
	return []int{}

}
