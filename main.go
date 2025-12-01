package main

import (
	"go-course/lessons"
)

func main() {
	user := lessons.User{
		ID:    1,
		Name:  "Иван Петров",
		Email: "ivan.petrov@example.com",
		Phone: "+7 999 123-45-67",
		Address: lessons.Address{
			Street:     "Улица Ленина",
			City:       "Москва",
			PostalCode: "101000",
		},
		Cart: []lessons.CartItem{
			{
				Product: lessons.Product{
					ID:          1,
					Name:        "Ноутбук",
					Description: "Мощный ноутбук для работы и игр",
					Price:       59990,
					Category:    "Электроника",
					Brand:       "Brand A",
					Rating:      4.5,
					Reviews:     120,
				},
				Quantity: 1,
			},
			{
				Product: lessons.Product{
					ID:          2,
					Name:        "Смартфон",
					Description: "Современный смартфон с отличной камерой",
					Price:       29990,
					Category:    "Электроника",
					Brand:       "Brand B",
					Rating:      4.7,
					Reviews:     200,
				},
				Quantity: 2,
			},
			{
				Product: lessons.Product{
					ID:          3,
					Name:        "Наушники",
					Description: "Беспроводные наушники с шумоподавлением",
					Price:       7990,
					Category:    "Аудио",
					Brand:       "Brand C",
					Rating:      4.3,
					Reviews:     80,
				},
				Quantity: 1,
			},
		},
	}

	lessons.PrintInfo(user)
}
