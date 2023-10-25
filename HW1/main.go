package main

import (
	"fmt"

	list "learn.go/datastructures"
)

func main() {
	//Тесты некоторого функционала
	example := list.New()
	for i := 0; i < 10; i++ {
		example.Append(i)
	}
	example.PrintVals()
	fmt.Print("\n")
	testslice := []int{1, 5, 7, 4, 3, 6}
	example = list.NewFromSlice(testslice)
	example.PrintVals()
}
