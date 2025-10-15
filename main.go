package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"strconv"
)

func main() {
	random := math.Floor((rand.Float64()*100)*10) / 10
	var isEven bool
	if math.Mod(random, 2) == 0 {
		isEven = true
	} else {
		isEven = false
	}

	fmt.Println("Исходное число: ", random)
	fmt.Println("Исходное число, увеличенное на 10%: ", strconv.FormatFloat(random*1.1, 'f', 5, 64))
	fmt.Println("Исходное число является четным: ", isEven)
	fmt.Println("Предпоследняя цифра целой части исходного числа: ", math.Floor(random/10))
}
