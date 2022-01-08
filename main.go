package main

import (
	"fmt"

	"github.com/AntonyIS/Golang-DSA/solutions"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := solutions.TwoSum(nums, target)
	fmt.Println(res)
}
