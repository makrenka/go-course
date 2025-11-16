package main

import (
	"go-course/lessons"
)

func main() {
	arr := [5]int{3, 8, 1, 8, 1}
	lessons.GenerateSecretCode(arr)
}
