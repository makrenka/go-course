package lessons

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const (
	MinPasswordLength = 4
	MinPasswordsCount = 1
	MaxPasswordsCount = 50
)

var (
	ErrPasswordLengthTooLow = errors.New("password length too low")
	ErrTooLowPasswordsCount = errors.New("too low passwords count")
	ErrTooBigPasswordsCount = errors.New("too big passwords count")
)

var (
	upperChars   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowerChars   = []rune("abcdefghijklmnopqrstuvwxyz")
	digitChars   = []rune("0123456789")
	specialChars = []rune("!@#$%^&*")
)

func GeneratePassword(length, count int) ([]string, error) {
	if length < MinPasswordLength {
		return nil, ErrPasswordLengthTooLow
	}
	if count < MinPasswordsCount {
		return nil, ErrTooLowPasswordsCount
	}

	generatedPasswords := make([]string, length)

	for i := 1; i <= length; i++ {
		upperInd, _ := rand.Int(rand.Reader, big.NewInt(int64(len(upperChars))))
		lowerInd, _ := rand.Int(rand.Reader, big.NewInt(int64(len(lowerChars))))
		digitInd, _ := rand.Int(rand.Reader, big.NewInt(int64(len(digitChars))))
		specialInd, _ := rand.Int(rand.Reader, big.NewInt(int64(len(specialChars))))

		password := string(upperChars[int(upperInd.Int64())]) +
			string(lowerChars[int(lowerInd.Int64())]) +
			string(digitChars[int(digitInd.Int64())]) +
			string(specialChars[int(specialInd.Int64())])

		generatedPasswords = append(generatedPasswords, password)
	}

	m := make(map[string]struct{}, len(generatedPasswords))
	res := make([]string, length)
	for i := 0; i < len(generatedPasswords); i++ {
		val := generatedPasswords[i]
		if _, ok := m[val]; ok {
			continue
		}
		m[val] = struct{}{}
		res = append(res, val)
	}

	return res, nil
}
