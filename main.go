package main

import (
	"go-course/lessons"
)

func main() {
	w, _ := lessons.NewWarrior("Король Артур")
	m, _ := lessons.NewMage("Мерлин")
	a, _ := lessons.NewArcher("Леголас")

	characters := []lessons.Character{w, m, a}
	lessons.Fight(characters)
}
