package main

import (
	"fmt"

	"github.com/AntonyIS/Golang-DSA/solutions"
)

func main() {
	prices := []int{8, 2, 7, 1, 5, 3, 6, 4}
	res := solutions.BestTimeToBuyStock(prices)
	fmt.Println(res)
}
