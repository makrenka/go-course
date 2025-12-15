package main

import (
	"errors"
	"fmt"
	"go-course/practices"
	"os"
)

func main() {
	if err := practices.HandlePayments(); err != nil {
		// Обработать ошибки, которые могут произойти при вызове HandlePayments
		var insufficientFundsErr *practices.InsufficientFundsError

		switch {
		case errors.Is(err, practices.ErrUnsupportedPaymentMethod):
			fmt.Println("Ошибка: неподдерживаемый метод платежа.")
		case errors.As(err, &insufficientFundsErr):
			fmt.Printf("Ошибка: недостаточно средств. %v\n", err)
		default:
			fmt.Printf("Неизвестная ошибка: %s.\n", err)
		}

		os.Exit(1)
	}

	fmt.Println("Платеж успешно обработан!")
}
