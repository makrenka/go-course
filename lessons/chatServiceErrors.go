package lessons

import (
	"errors"
	"fmt"
	"log"
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
		if errors.As(err, &getChatError) {
			if getChatError.Code == 24 {
				log.Fatalf("Ошибка запроса: %v", err)
			} else {
				log.Fatalf("Инфраструктурная ошибка: %v", err)
			}
		} else {
			log.Fatalf("Неизвестная ошибка: %v", err)
		}
	}

	for _, v := range chat.Messages {
		fmt.Printf("%s: %s", v.From, v.Message)
	}
}
