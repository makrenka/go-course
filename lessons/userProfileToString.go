package lessons

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"
)

func UserProfileToString(name string, age int) (string, error) {
	message := fmt.Sprintf("Имя человека: %s, возраст: %d.", strings.TrimSpace(name), age)
	if name == "" {
		return "", errors.New("empty name")
	}
	if age < 0 {
		return "", errors.New("negative age")
	}
	if strings.TrimSpace(name) == "" {
		return "", errors.New("name cannot contain only spaces")
	}
	return message, nil
}

func GenerateCompliment(name string) string {
	randomNum := rand.IntN(3)
	var message string

	switch {
	case randomNum == 0:
		message = "Ты великолепен, " + name + "!"
	case randomNum == 1:
		message = "У тебя потрясающая улыбка, " + name + "!"
	case randomNum == 2:
		message = "Ты вдохновляешь, " + name + "!"
	default:
		fmt.Println("Что-то пошло не так")
	}

	return message
}
