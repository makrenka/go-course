package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func main() {
	random := math.Floor((rand.Float64()*100)*10) / 10
	fmt.Println("Исходное число: ", random)
	fmt.Println("Исходное число, увеличенное на 10%: ", math.Round(random*1.1*100000)/100000)
	fmt.Println("Исходное число является четным: ", math.Mod(random, 2))
}
