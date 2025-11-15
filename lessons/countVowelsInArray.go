package lessons

import (
	"fmt"
	"strings"
)

func CountVowelsInArray(arr [3]string) {
	for _, v := range arr {
		vowels := "аеёиоуыэюя"
		var count int
		for _, char := range strings.ToLower(v) {
			if strings.ContainsRune(vowels, char) {
				count++
			}
		}
		fmt.Print(count, " ")
	}
}
