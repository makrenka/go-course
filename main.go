package main

import (
	"fmt"
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

	printNumberInfo(0)
}

func printNumberInfo(num int) {
	if math.Abs(float64(num)) != float64(num) {
		fmt.Printf("Число %d отрицательное.\n", num)
	} else if num == 0 {
		fmt.Printf("Число равно %d.\n", num)
	} else {
		fmt.Printf("Число %d положительное.\n", num)

	}

	if num%2 != 0 {
		fmt.Printf("Число %d нечетное.\n", num)
	} else {
		fmt.Printf("Число %d четное.\n", num)
	}

	sqrt := math.Sqrt(float64(num))
	if math.Abs(float64(num)) == float64(num) && num%2 != 0 {
		fmt.Printf("Квадратный корень числа %d не является целым числом и равен %.5f.\n", num, sqrt)
	} else if math.Abs(float64(num)) == float64(num) && num%2 == 0 && num != 0 {
		fmt.Printf("Квадратный корень числа %d является целым числом и равен %d.\n", num, int(sqrt))
	}
}

// ---------------------------------------------------------------------------------------
// На основе полученного значения ИМТ, программа должна определить и вывести соответствующую категорию:
// Недостаточный вес: ИМТ < 18.5
// Нормальный вес: 18.5 ≤ ИМТ < 25
// Избыточный вес: 25 ≤ ИМТ < 30
// Ожирение: ИМТ ≥ 30
// func bodyMassIndex() {
// 	var height float64
// 	var weight float64

// 	fmt.Printf("Введите ваш вес (кг): ")
// 	_, errWeight := fmt.Scan(&weight)
// 	if errWeight != nil || weight <= 0 {
// 		log.Fatal("Неверно задан вес")
// 	}

// 	fmt.Printf("Введите ваш рост (см): ")
// 	_, errHeight := fmt.Scan(&height)
// 	if errHeight != nil || height <= 0 {
// 		log.Fatal("Неверно задан рост")
// 	}

// 	bodyIndex := weight / math.Pow((height/100), 2)
// 	fmt.Printf("Ваш ИМТ: %.2f\n", bodyIndex)

// 	var message string

// 	switch {
// 	case bodyIndex < 18.5:
// 		message = "Недостаточный вес"
// 	case bodyIndex < 25:
// 		message = "Нормальный вес"
// 	case bodyIndex < 30:
// 		message = "Избыточный вес"
// 	case bodyIndex >= 30:
// 		message = "Ожирение"
// 	default:
// 		message = "Так не бывает"
// 	}

// 	fmt.Println("Категория:", message)
// }

// -----------------------------------------------------------------------------------------
// Поиск по названию товара

// func searchByProductName() {
// 	const (
// 		keyboard   = "Клавиатура JZ9: 19200"
// 		headphones = "Наушники N45: 9600"
// 		smartphone = "Смартфон S10: 55000"
// 	)

// 	fmt.Printf("Введите название товара: ")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	if err := scanner.Err(); err != nil {
// 		log.Fatalf("Ошибка, %s", err)
// 	}

// 	input := strings.ToLower(scanner.Text())
// 	var message string

// 	switch {
// 	case strings.Contains(strings.ToLower(keyboard), input):
// 		message = keyboard
// 	case strings.Contains(strings.ToLower(headphones), input):
// 		message = headphones
// 	case strings.Contains(strings.ToLower(smartphone), input):
// 		message = smartphone
// 	default:
// 		fmt.Printf("Товар %q не найден", input)
// 	}

// 	fmt.Println(message)
// }
