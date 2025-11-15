package lessons

import "fmt"

func IsPalindrome(arr [10]int) {
	for i := 0; i < 5; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			fmt.Println("Не палиндром!")
			return
		}
	}

	fmt.Println("Это палиндром!")
}
