package lessons

import (
	"fmt"
	"maps"
)

func RemoveSlicesBySum(m map[string][]int) {
	maps.DeleteFunc(m, func(key string, val []int) bool {
		sum := 0
		for _, v := range val {
			sum += v
		}
		return sum > 6
	})
	fmt.Println(m)
}
