package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	// random := math.Floor((rand.Float64()*100)*10) / 10
	// var isEven bool
	// if math.Mod(random, 2) == 0 {
	// 	isEven = true
	// } else {
	// 	isEven = false
	// }

	// fmt.Println("Исходное число: ", random)
	// fmt.Println("Исходное число, увеличенное на 10%: ", strconv.FormatFloat(random*1.1, 'f', 5, 64))
	// fmt.Println("Исходное число является четным: ", isEven)
	// fmt.Println("Предпоследняя цифра целой части исходного числа: ", math.Floor(random/10))

	bodyMassIndex()
}

// На основе полученного значения ИМТ, программа должна определить и вывести соответствующую категорию:
// Недостаточный вес: ИМТ < 18.5
// Нормальный вес: 18.5 ≤ ИМТ < 25
// Избыточный вес: 25 ≤ ИМТ < 30
// Ожирение: ИМТ ≥ 30
func bodyMassIndex() {
	var height float64
	var weight float64

	fmt.Printf("Введите ваш вес (кг): ")
	_, errWeight := fmt.Scan(&weight)
	if errWeight != nil || weight <= 0 {
		log.Fatal("Неверно задан вес")
	}

	fmt.Printf("Введите ваш рост (см): ")
	_, errHeight := fmt.Scan(&height)
	if errHeight != nil || height <= 0 {
		log.Fatal("Неверно задан рост")
	}

	bodyIndex := weight / math.Pow((height/100), 2)
	fmt.Printf("Ваш ИМТ: %.2f\n", bodyIndex)

	var message string

	switch {
	case bodyIndex < 18.5:
		message = "Недостаточный вес"
	case bodyIndex < 25:
		message = "Нормальный вес"
	case bodyIndex < 30:
		message = "Избыточный вес"
	case bodyIndex >= 30:
		message = "Ожирение"
	default:
		message = "Так не бывает"
	}

	fmt.Println("Категория:", message)
}
