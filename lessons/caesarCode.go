package lessons

import (
	"unicode"
)

func CaesarCode(text string, shift int32, encode bool) string {
	alphabet := []rune("абвгдеёжзийклмнопрстуфхцчшщъыьэюя")
	alphabetUpper := []rune("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ")
	res := []rune(text)

	if !encode {
		shift = -shift
	}

	for i, r := range res {
		switch {
		case unicode.IsLower(r):
			for j, v := range alphabet {
				if r == v {
					newIndex := (j + int(shift)) % len(alphabet)
					if newIndex < 0 {
						newIndex += len(alphabet)
					}
					res[i] = alphabet[newIndex]
					break
				}
			}
		case unicode.IsUpper(r):
			for j, v := range alphabetUpper {
				if r == v {
					newIndex := (j + int(shift)) % len(alphabetUpper)
					if newIndex < 0 {
						newIndex += len(alphabetUpper)
					}
					res[i] = alphabetUpper[newIndex]
					break
				}
			}
		default:
			continue
		}

	}

	return string(res)
}
