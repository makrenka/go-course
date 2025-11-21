package main

import (
	"fmt"
	"go-course/lessons"
)

func main() {
	str := "Зашифруй меня!"
	encodedStr := lessons.CaesarCode(str, 5, true)
	fmt.Println(encodedStr)

	decodedStr := lessons.CaesarCode(encodedStr, 5, false)
	fmt.Println(decodedStr)
}
