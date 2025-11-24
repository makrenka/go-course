package main

import (
	"go-course/lessons"
)

func main() {
	// res, _ := lessons.CreateSlice(6)
	lessons.SortByParity([]int{-7, -5, 6, -10, 2, -9})
}
