package main

import (
	"go-course/lessons"
)

func main() {
	slice1 := []int{1, 2, 3, 3, 5}
	slice2 := []int{3, 3, 3, 4, 5, 6, 7}
	lessons.IntersectSlices(slice1, slice2)
}
