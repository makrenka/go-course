package main

import (
	"go-course/lessons"
)

func main() {
	slice1 := [][]int{
		{3, 1, 4, 1},
		{2, 2, 2},
		{5, 0, 6, 3, -8, 1},
		{4, 6, 8, 2},
	}
	lessons.MagicSort(slice1)
}
