package lessons

import (
	"errors"
	"fmt"
)

// Данные
type Message struct {
	From    string
	Message string
}

type Chat struct {
	ID       int
	Messages []Message
}

// Ошибка базы данных
type DatabaseError struct {
	Message string
	Code    int
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("database error with code %d: %s", e.Code, e.Message)
}

// Сервис для работы
type Service struct {
	db interface {
		GetChatByIDWithMessages(id int) (*Chat, error)
	}
}

func (w Service) PrintChat(id int) {
	var getChatError *DatabaseError
	chat, err := w.db.GetChatByIDWithMessages(id)
	if err != nil {
		switch {
		case errors.As(err, &getChatError) && getChatError.Code == 24:
			fmt.Printf("Ошибка запроса: %s", getChatError)
		case errors.As(err, &getChatError):
			fmt.Printf("Инфраструктурная ошибка: %s", getChatError)
		default:
			fmt.Printf("Неизвестная ошибка: %s", err)
		}
		return
	}

	for _, v := range chat.Messages {
		fmt.Printf("%s: %s", v.From, v.Message)
	}
}
