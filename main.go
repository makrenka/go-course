package main

import (
	"go-course/lessons"
)

func main() {
	slice := [][]int{
		{1, 2, 3},
		{},
		{7, 8, 9, 22, 48},
		{10, 11},
	}
	lessons.ReplaceEvenOnEvenIndices(slice)
}
