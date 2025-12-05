package lessons

import (
	"errors"
	"fmt"
)

type Warrior struct {
	Name string
}

func NewWarrior(name string) (*Warrior, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	return &Warrior{
		Name: name,
	}, nil
}

func (w Warrior) Attack() string {
	return fmt.Sprintf("Воин %s бьет мечом.", w.Name)
}

type Mage struct {
	Name string
}

func NewMage(name string) (*Mage, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	return &Mage{
		Name: name,
	}, nil
}

func (m Mage) Attack() string {
	return fmt.Sprintf("Маг %s колдует огненный шар.", m.Name)
}

type Archer struct {
	Name string
}

func NewArcher(name string) (*Archer, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	return &Archer{
		Name: name,
	}, nil
}

func (a Archer) Attack() string {
	return fmt.Sprintf("Лучник %s выпускает град стрел.", a.Name)
}

type Character interface {
	Attack() string
}

func Fight(characters []Character) {
	for _, c := range characters {
		fmt.Println(c.Attack())
	}
}
