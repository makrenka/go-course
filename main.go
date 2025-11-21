package main

import (
	"go-course/lessons"
)

func main() {
	res, _ := lessons.CreateSlice(5)
	lessons.FilterSlice(res)
}
