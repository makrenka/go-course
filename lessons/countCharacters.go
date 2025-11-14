package lessons

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"unicode"
)

func getInput() (string, error) {
	fmt.Println("Введите текст")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if input == "" {
		return "", errors.New("неверный ввод")
	}

	return input, nil
}

func CountCharacters() (letters, digits, spaces, punctuation int) {
	text, err := getInput()

	if err != nil {
		fmt.Println("Неверный ввод:", err)
		return
	}

	for _, char := range text {
		switch {
		case unicode.IsDigit(char):
			digits++
		case unicode.IsLetter(char):
			letters++
		case unicode.IsPunct(char):
			punctuation++
		case unicode.IsSpace(char):
			spaces++
		}
	}

	displayResults(letters, digits, spaces, punctuation)

	return letters, digits, spaces, punctuation
}

func displayResults(letters, digits, spaces, punctuation int) {
	fmt.Println("Количество букв:", letters)
	fmt.Println("Количество цифр:", digits)
	fmt.Println("Количество пробелов:", spaces)
	fmt.Println("Количество знаков препинания:", punctuation)
}
