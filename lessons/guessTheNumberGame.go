package lessons

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func getRandom() int {
	fmt.Println("Компьютер загадал случайное число от 1 до 100 включительно. Угадайте его!")
	return rand.IntN(100) + 1
}

func guessNumber(num, random int) int {
	if num > random {
		return -1
	}
	if num < random {
		return 1
	}
	return 0
}

func GetUserNmuber() *int {
	guesses := 1
	random := getRandom()
	fmt.Println("Компьютер загадал:", random)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Ваше предположение (либо, для завершения, введите слово выход):")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "выход" {
			fmt.Println("Выход из программы...")
			return nil
		}

		number, err := strconv.Atoi(input)

		if err != nil || number < 1 || number > 100 {
			fmt.Println("Ошибка: введите целое число от 1 до 100 или \"выход\"")
			continue
		}

		result := guessNumber(number, random)
		if result == -1 {
			fmt.Println("Загаданное число меньше.")
			guesses++
			continue
		} else if result == 1 {
			fmt.Println("Загаданное число больше.")
			guesses++
			continue
		} else if result == 0 {
			fmt.Printf("Правильно! Вы угадали число с %d попытки.\n", guesses)
			fmt.Println("Хотите сыграть еще раз? (если хотите, напишите слово да):")
		}

		for {
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(strings.ToLower(answer))
			switch answer {
			case "да":
				return GetUserNmuber()
			case "нет":
				fmt.Println("Спасибо за игру! До свидания!")
				return nil
			default:
				fmt.Println("Введите: да или нет")
				continue
			}
		}
	}
}
