package main

import (
	"github.com/AntonyIS/Golang-DSA/solutions"
)

func main() {
	linkedList := solutions.LinkedList{}

	linkedList.Prepend(10)
	linkedList.Prepend(20)
	linkedList.Prepend(30)
	linkedList.Prepend(40)
	linkedList.Prepend(50)
	linkedList.Append(60)
	linkedList.Append(70)
	linkedList.Append(80)
	linkedList.Append(90)
	linkedList.Append(100)

	linkedList.PrintList()
	linkedList.Insert(22, 0)
	linkedList.PrintList()

}
