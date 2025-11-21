package lessons

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func CreateSlice(n int) ([]int, error) {
	if n < 0 {
		fmt.Println("число должно быть больше нуля")
		return nil, errors.New("число должно быть больше нуля")
	}

	res := make([]int, n)
	for i := range res {
		res[i] = rand.IntN(20) - 10
	}
	fmt.Println(res)
	return res, nil
}

func FilterSlice(numbers []int) []int {
	var res []int
	for i := range numbers {
		if i != 0 &&
			numbers[i] < numbers[i-1] &&
			(numbers[i]%2 == 0 || numbers[i]%5 == 0 || numbers[i]%6 == 0 || numbers[i]%9 == 0) {
			res = append(res, numbers[i])
		}
	}
	fmt.Println(res)
	return res
}
