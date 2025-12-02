package lessons

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

type User1 struct {
	FirstName         string
	LastName          string
	BirthYear         int
	FavoriteLanguages []string
}

func (u User1) SecretIdentity() string {
	firstLetter, _ := utf8.DecodeRuneInString(u.FirstName)
	secondLetter, _ := utf8.DecodeRuneInString(u.LastName)
	return string(firstLetter) + string(secondLetter) + strconv.Itoa(rand.IntN(100)+1)
}

func (u User1) Age() int {
	return time.Now().Year() - u.BirthYear
}

func (u *User1) AddFavoriteLanguage(language string) error {
	if len(language) == 0 {
		return errors.New("empty language name")
	}

	if !slices.Contains(u.FavoriteLanguages, language) {
		u.FavoriteLanguages = append(u.FavoriteLanguages, language)
		return nil
	}

	return errors.New("duplicate")
}

func (u *User1) RemoveFavoriteLanguage(language string) error {
	if !slices.Contains(u.FavoriteLanguages, language) {
		return errors.New("not found")
	}

	ind := slices.Index(u.FavoriteLanguages, language)
	u.FavoriteLanguages = append(u.FavoriteLanguages[:ind], u.FavoriteLanguages[ind+1:]...)
	return nil
}

func (u User1) IsProgrammingLanguageFavorite(language string) bool {
	return slices.Contains(u.FavoriteLanguages, language)
}

func (u User1) RandomFavoriteLanguage() (string, error) {
	if len(u.FavoriteLanguages) != 0 {
		randInd := rand.IntN(len(u.FavoriteLanguages))
		return u.FavoriteLanguages[randInd], nil
	}
	return "", errors.New("no options")
}

func (u User1) GenerateProfile() string {
	favoriteLangStr := strings.Join(u.FavoriteLanguages, ", ")
	res := fmt.Sprintf("Имя: %s.\nФамилия: %s.\nВозраст: %d.\nСписок любимых языков программирования: [%s].", u.FirstName, u.LastName, u.Age(), favoriteLangStr)
	return res
}

func (u *User1) UpdateName(firstName, lastName string) error {
	if firstName == "" || lastName == "" {
		return errors.New("empty data")
	}

	firstRuneFirstName, _ := utf8.DecodeRuneInString(firstName)
	firstRuneLastName, _ := utf8.DecodeRuneInString(lastName)

	if unicode.IsUpper(firstRuneFirstName) && unicode.IsUpper(firstRuneLastName) {
		u.FirstName = firstName
		u.LastName = lastName
	} else {
		return errors.New("invalid data")
	}

	return nil
}
