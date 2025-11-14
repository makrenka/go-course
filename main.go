package main

import (
	"go-course/lessons"
)

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 4, 3, 2, 1, 1}
	lessons.IsPalindrome(arr)
}
