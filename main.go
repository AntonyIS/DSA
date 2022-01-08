package main

import (
	"fmt"

	"github.com/AntonyIS/Golang-DSA/solutions"
)

func main() {
	nums := []int{1, 2, 3}
	res := solutions.ContainsDuplicates(nums)
	fmt.Println(res)
}
