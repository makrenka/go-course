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
		var getInsufficientFundsError *practices.InsufficientFundsError

		switch {
		case errors.Is(err, practices.ErrUnsupportedPaymentMethod):
			fmt.Println("Ошибка: неподдерживаемый метод платежа.")
		case errors.As(err, &getInsufficientFundsError):
			fmt.Printf("Ошибка: недостаточно средств. requested amount: %.2f, maximum allowed amount: %.2f\n",
				getInsufficientFundsError.RequestedAmount,
				getInsufficientFundsError.MaxAmount)
		default:
			fmt.Printf("Неизвестная ошибка: %s.", err)
		}

		os.Exit(1)
	}

	fmt.Println("Платеж успешно обработан!")
}
