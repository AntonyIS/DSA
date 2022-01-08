package solutions

func Merge_sort(list []int) []int {
	/*
		divide: divid a list into sublist
		conquer : sort the sublist in the previous step
		combine : merges the sublist
	*/

	if len(list) <= 1 {
		return list
	} else {
		left_half, right_half := split(list)
		left := Merge_sort(left_half)
		right := Merge_sort(right_half)
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
