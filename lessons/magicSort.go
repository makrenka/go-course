package lessons

import (
	"slices"
)

// маё
// func MagicSort(s [][]int) {
// 	slices.SortStableFunc(s, func(a, b []int) int {
// 		return sum(a) - sum(b)
// 	})

// 	for i, slice := range s {
// 		var evens []int
// 		var odds []int
// 		var zeros []int
// 		for _, v := range slice {
// 			switch {
// 			case v != 0 && v%2 == 0:
// 				evens = append(evens, v)
// 			case v%2 != 0:
// 				odds = append(odds, v)
// 			default:
// 				zeros = append(zeros, v)
// 			}
// 		}
// 		slices.SortFunc(evens, func(a, b int) int {
// 			return b - a
// 		})
// 		slices.SortFunc(odds, func(a, b int) int {
// 			return b - a
// 		})

// 		newSlice := append(zeros, evens...)
// 		newSlice = append(newSlice, odds...)
// 		s[i] = newSlice
// 	}

// 	fmt.Println(s)
// }

// func sum(slice []int) int {
// 	total := 0
// 	for _, v := range slice {
// 		total += v
// 	}
// 	return total
// }

// Лепшае
func isEven(num int) bool {
	return num%2 == 0
}

func compSliceSum(sl1, sl2 []int) int {
	sum1, sum2 := 0, 0
	for _, v := range sl1 {
		sum1 += v
	}
	for _, v := range sl2 {
		sum2 += v
	}
	return sum1 - sum2
}

func compSliceEvens(a, b int) (res int) {
	if a == 0 {
		res = -1
	} else if b == 0 {
		res = 1
	} else if isEven(a) && !isEven(b) {
		res = -1
	} else if isEven(b) && !isEven(a) {
		return 1
	} else {
		res = b - a
	}
	return
}

func MagicSort(sl [][]int) {
	for _, subSlice := range sl {
		slices.SortFunc(subSlice, compSliceEvens)
	}
	slices.SortStableFunc(sl, compSliceSum)
}
