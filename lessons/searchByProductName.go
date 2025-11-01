package lessons

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Поиск по названию товара

func SearchByProductName() {
	const (
		keyboard   = "Клавиатура JZ9: 19200"
		headphones = "Наушники N45: 9600"
		smartphone = "Смартфон S10: 55000"
	)

	fmt.Printf("Введите название товара: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка, %s", err)
	}

	input := strings.ToLower(scanner.Text())
	var message string

	switch {
	case strings.Contains(strings.ToLower(keyboard), input):
		message = keyboard
	case strings.Contains(strings.ToLower(headphones), input):
		message = headphones
	case strings.Contains(strings.ToLower(smartphone), input):
		message = smartphone
	default:
		fmt.Printf("Товар %q не найден", input)
	}

	fmt.Println(message)
}
