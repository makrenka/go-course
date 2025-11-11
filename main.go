package main

import (
	"fmt"
	"go-course/lessons"
	"math/rand/v2"
	"os"
)

func main() {
	for range 100 {
		random := rand.IntN(100) + 1
		result := lessons.Play()
		if result != random {
			fmt.Printf("Неверный ответ. Было загадано число %d, а в ответе получили число %d", random, result)
			os.Exit(-1)
		}
	}
}
