package lessons

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"slices"
	"strconv"
	"time"
	"unicode"
)

type User1 struct {
	FirstName         string
	LastName          string
	BirthYear         int
	FavoriteLanguages []string
}

func (u User1) SecretIdentity() string {
	return rune(u.FirstName[0]) + rune(u.LastName[0]) + strconv.Itoa(rand.IntN(100)+1)
}

func (u User1) Age() int {
	return time.Now().Year() - u.BirthYear
}

func (u *User1) AddFavoriteLanguage(language string) error {
	if slices.Contains(u.FavoriteLanguages, language) {
		u.FavoriteLanguages = append(u.FavoriteLanguages, language)
	}
	return errors.New("duplicate")
}

func (u *User1) RemoveFavoriteLanguage(language string) error {
	ind := slices.Index(u.FavoriteLanguages, language)
	u.FavoriteLanguages = append(u.FavoriteLanguages[:ind], u.FavoriteLanguages[ind+1:]...)
	return errors.New("not found")
}

func (u User1) IsProgrammingLanguageFavorite(language string) bool {
	return slices.Contains(u.FavoriteLanguages, language)
}

func (u User1) RandomFavoriteLanguage() (string, error) {
	if len(u.FavoriteLanguages) != 0 {
		randInd := rand.IntN(len(u.FavoriteLanguages))
		return u.FavoriteLanguages[randInd], nil
	}
	return "", errors.New("not realized")
}

func (u User1) GenerateProfile() string {
	res := fmt.Sprintf("Имя: %s.\nФамилия: %s.\nВозраст: %d.\nСписок любимых языков программирования: %v.", u.FirstName, u.LastName, u.Age(), u.FavoriteLanguages)
	return res
}

func (u *User1) UpdateName(firstName, lastName string) error {
	if firstName == "" || lastName == "" {
		return errors.New("empty data")
	}
	if unicode.IsLower(rune(firstName[0])) || unicode.IsLower(rune(lastName[0])) {
		return errors.New("invalid data")
	}
	u.FirstName = firstName
	u.LastName = lastName
	return nil
}
