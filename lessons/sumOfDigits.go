package lessons

import (
	"strconv"
	"strings"
)

func SumOfDigits(n int) int {
	if n < 0 {
		n = -n
	}

	arr := strings.Split(strconv.Itoa(n), "")

	if len(arr) == 1 {
		str := strings.Join(arr, "")
		num, _ := strconv.Atoi(str)
		return num
	}

	num, _ := strconv.Atoi(arr[0])
	arr = arr[1:]
	newStr := strings.Join(arr, "")
	newNum, _ := strconv.Atoi(newStr)

	return num + SumOfDigits(newNum)
}
