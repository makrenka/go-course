package lessons

import "fmt"

func IsPalindrome(arr [10]int) {
	var reversed [10]int

	for j := range reversed {
		reversed[j] = arr[len(arr)-j-1]
	}

	var count int
	for i := range arr {
		if arr[i] == reversed[i] {
			count++
		}
	}

	if count == 10 {
		fmt.Println("Это палиндром!")
	} else {
		fmt.Println("Не палиндром!")
	}
}
