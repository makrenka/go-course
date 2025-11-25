package main

import (
	"go-course/lessons"
)

func main() {
	m1 := map[string][]int{
		"a": {1, 2, 3},
		// "b": {1, 2, 3},
	}
	m2 := map[string][]int{
		"a": {},
		// "b": {3, 2, 1},
	}
	lessons.CompareMaxValues(m1, m2)
}
