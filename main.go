package main

import (
	"fmt"

	"github.com/AntonyIS/Golang-DSA/solutions"
)

func main() {
	sorted_list := []int{1, 2, 4, 5, 6, 7, 8, 9}
	unsorted_list := []int{9, 2, 6, 1, 3, 7, 5}

	res := solutions.Merge_sort(unsorted_list)
	fmt.Println(res)
	// linear_search
	ok, err := solutions.Linear_search(sorted_list, 9)
	fmt.Println("LINEAR SEARCH")
	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println(ok)
	}
	// binary_search
	fmt.Println("BINARY SEARCH")
	pass, err := solutions.Binary_search(sorted_list, 9)

	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println(pass)
	}

	// binary_search
	fmt.Println("RECURSIVE BINARY  SEARCH")
	check, err := solutions.Recursive_binary_search(sorted_list, 9)

	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println(check)
	}
}
