package main

import (
	"fmt"
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

func MagicSort(s [][]int) [][]int {
	for _, slice := range s {
		sum := 0
		for _, v := range slice {
			sum += v
		}
		slice = append(slice, sum)
	}

	fmt.Println(s)
	return s
}
