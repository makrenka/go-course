package main

import (
	"fmt"
	"go-course/practices"
)

func main() {
	tm := practices.NewTagManager()

	// Добавление тегов
	if err := tm.AddTag("golang"); err != nil {
		fmt.Println(err)
	}
	if err := tm.AddTag("programming"); err != nil {
		fmt.Println(err)
	}
	if err := tm.AddTag("golang"); err != nil {
		fmt.Println(err) // Ошибка, тег уже существует
	}

	// Проверка существования тегов
	fmt.Println("Тег 'golang' существует:", tm.TagExists("golang")) // true
	fmt.Println("Тег 'python' существует:", tm.TagExists("python")) // false

	// Список тегов
	fmt.Println("Current tags:", tm.ListTags()) // [golang programming]

	// Удаление тегов
	if err := tm.RemoveTag("golang"); err != nil {
		fmt.Println(err)
	}
	if err := tm.RemoveTag("golang"); err != nil {
		fmt.Println(err) // Ошибка, тег не существует
	}

	// Список тегов после удаления
	fmt.Println("Current tags after removal:", tm.ListTags()) // [programming]
}
