package lessons

import (
	"fmt"
	"strings"
)

// User представляет пользователя
type User struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Address Address
	Cart    []CartItem
}

// Address представляет адрес пользователя
type Address struct {
	Street     string
	City       string
	PostalCode string
}

// CartItem представляет элемент в корзине
type CartItem struct {
	Product  Product
	Quantity int
}

// Product представляет продукт в корзине
type Product struct {
	ID          int
	Name        string
	Description string
	Price       int
	Category    string
	Brand       string
	Rating      float64
	Reviews     int
}

func PrintInfo(user User) {
	fmt.Printf("Покупатель %s. Телефон: %s. Адрес: г. %s, %s.\n", user.Name, user.Phone, user.Address.City, user.Address.Street)

	if isElectronicBuyer(user.Cart) {
		fmt.Println("Пользователь является покупателем электроники.")
	} else {
		fmt.Println("Пользователь не является покупателем электроники.")
	}

	fmt.Printf("Товары в корзине, где цена 10000 и более: %s.\n", highPriceItems(user.Cart))
	fmt.Printf("Общая сумма покупки: %d руб.", totalAmount(user.Cart))
}

func isElectronicBuyer(cart []CartItem) bool {
	for _, v := range cart {
		if v.Product.Category == "Электроника" {
			return true
		}
	}
	return false
}

func highPriceItems(cart []CartItem) string {
	items := make([]string, 0)
	for _, v := range cart {
		if v.Product.Price > 10000 {
			items = append(items, v.Product.Name)
		}
	}

	if len(items) == 0 {
		return "отсутствуют"
	}

	return strings.Join(items, ", ")
}

func totalAmount(cart []CartItem) int {
	res := 0
	for _, v := range cart {
		res += v.Product.Price * v.Quantity
	}
	return res
}
