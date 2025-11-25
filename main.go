package main

import (
	"go-course/lessons"
)

func main() {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	m2 := map[string]int{
		"b": 3,
		"c": 4,
		"d": 5,
	}
	lessons.MergeMaps(m1, m2)
}

func MergeMaps(m1, m2 map[string]int) map[string]int {
	res := make(map[string]int)
	for key1, v1 := range m1 {
		res[key1] = v1
		for key2, v2 := range m2 {
			if key2 == key1 {
				res[key2] = v2 + v1
			} else {
				res[key2] = v2
			}
		}
	}
	return res
}
