package main

import (
	"go-course/lessons"
)

func main() {
	slice1 := []int{1, 1, 3, 4}
	slice2 := []int{1, 3, 4, 5}
	lessons.IntersectSlices(slice1, slice2)
}
