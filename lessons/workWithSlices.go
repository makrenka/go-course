package lessons

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"slices"
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

func MaxSumWithNegative(numbers []int, k int) []int {
	maxSumSlice := make([]int, k)
	var maxSum int
	for i := range numbers {
		var subSlice []int
		if i+k > len(numbers) {
			break
		}
		if i+k == len(numbers) {
			subSlice = numbers[i:]
		}
		if i+k < len(numbers) {
			subSlice = numbers[i : i+k]
		}
		var negativeSlice []int
		sum := 0
		for _, v := range subSlice {
			if v < 0 {
				negativeSlice = append(negativeSlice, v)
			}
		}
		if len(negativeSlice) > 0 {
			for _, v := range subSlice {
				sum += v
			}
		}
		if sum > maxSum || i == 0 {
			maxSum = sum
			copy(maxSumSlice, subSlice)
		}
	}
	fmt.Println(maxSumSlice)
	return maxSumSlice
}

func SortByParity(numbers []int) []int {
	slices.SortFunc(numbers, compareSliceNums)
	fmt.Println(numbers)
	return numbers
}

func isEvenNum(num int) bool {
	return num%2 == 0
}

func compareSliceNums(a, b int) (res int) {
	if isEvenNum(a) && !isEvenNum(b) {
		res = -1
	}
	if !isEvenNum(a) && isEvenNum(b) {
		res = 1
	}
	if isEvenNum(a) && isEvenNum(b) {
		res = b - a
	}
	if !isEvenNum(a) && !isEvenNum(b) {
		res = a - b
	}
	return
}
