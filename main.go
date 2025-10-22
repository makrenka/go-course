package main

// "errors"

// "strings"
// "math/rand/v2"

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

	add := adder(10)

	add(10)
}

func adder(n int) func(x int) int {
	res := n

	return func(x int) int {
		res += x
		return res
	}
}

// func calculate(n1 float64, n2 float64, action string) (float64, error) {
// 	if action == "add" {
// 		return n1 + n2, nil
// 	}
// 	if action == "subtract" {
// 		return n1 - n2, nil
// 	}
// 	if action == "multiply" {
// 		return n1 * n2, nil
// 	}
// 	if n2 == 0 {
// 		return 0, errors.New("division by zero")
// 	}
// 	if action == "divide" {
// 		return n1 / n2, nil
// 	}

// 	return 0, errors.New("unknown operation")
// }

// func UserProfileToString(name string, age int) (string, error) {
// 	message := fmt.Sprintf("Имя человека: %s, возраст: %d.", strings.TrimSpace(name), age)
// 	if name == "" {
// 		return "", errors.New("empty name")
// 	}
// 	if age < 0 {
// 		return "", errors.New("negative age")
// 	}
// 	if strings.TrimSpace(name) == "" {
// 		return "", errors.New("name cannot contain only spaces")
// 	}
// 	return message, nil
// }

// func generateCompliment(name string) string {
// 	randomNum := rand.IntN(3)
// 	var message string

// 	switch {
// 	case randomNum == 0:
// 		message = "Ты великолепен, " + name + "!"
// 	case randomNum == 1:
// 		message = "У тебя потрясающая улыбка, " + name + "!"
// 	case randomNum == 2:
// 		message = "Ты вдохновляешь, " + name + "!"
// 	default:
// 		fmt.Println("Что-то пошло не так")
// 	}

// 	return message
// }

// func printNumberInfo(num int) {
// 	if math.Abs(float64(num)) != float64(num) {
// 		fmt.Printf("Число %d отрицательное.\n", num)
// 	} else if num == 0 {
// 		fmt.Printf("Число равно %d.\n", num)
// 	} else {
// 		fmt.Printf("Число %d положительное.\n", num)

// 	}

// 	if num%2 != 0 {
// 		fmt.Printf("Число %d нечетное.\n", num)
// 	} else {
// 		fmt.Printf("Число %d четное.\n", num)
// 	}

// 	sqrt := math.Sqrt(float64(num))
// 	if math.Abs(float64(num)) == float64(num) && num%2 != 0 {
// 		fmt.Printf("Квадратный корень числа %d не является целым числом и равен %.5f.\n", num, sqrt)
// 	} else if math.Abs(float64(num)) == float64(num) && num%2 == 0 && num != 0 {
// 		fmt.Printf("Квадратный корень числа %d является целым числом и равен %d.\n", num, int(sqrt))
// 	}
// }

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
