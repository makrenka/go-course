package lessons

import (
	"fmt"
	// "slices"
)

func MagicSort(s [][]int) [][]int {
	for i, slice := range s {
		sum := 0
		for _, v := range slice {
			sum += v
		}
		s[i] = append(s[i], sum)
	}
	// s = slices.SortedStableFunc(s, func(i1, i2 []int) int {

	// 	return s[]
	// })

	fmt.Println(s)
	return s
}
