package main

import (
	"go-course/lessons"
)

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	lessons.SumNeighbors(arr)
}
