package main

import (
	"fmt"
)

func main() {
	hour := 25

	switch {
	case hour > 6 && hour <= 12:
		fmt.Printf("Сейчас %dч. - утро", hour)
	case hour > 12 && hour <= 18:
		fmt.Printf("Сейчас %dч. - день", hour)
	case hour > 18 && hour <= 23:
		fmt.Printf("Сейчас %dч. - вечер", hour)
	case hour > 23 && hour <= 24 && hour >= 0 && hour <= 6:
		fmt.Printf("Сейчас %dч. - ночь", hour)
	default:
		fmt.Printf("Неверно задано время")
	}
}
