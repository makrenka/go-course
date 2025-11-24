package main

import (
	"go-course/lessons"
)

func main() {
	m := map[string]int{
		"banana":     2,
		"apple":      1,
		"grapefruit": 3,
		"cherry":     1,
	}
	invertedMap := lessons.InvertMap(m)
	lessons.PrintMap(invertedMap)
}
