package lessons

import (
	"fmt"
	"maps"
	"slices"
)

func CompareMaxValues(m1, m2 map[string][]int) bool {
	res := maps.EqualFunc(m1, m2, func(sl1, sl2 []int) bool {
		if len(sl1) == 0 && len(sl2) == 0 {
			return true
		}
		if len(sl1) != len(sl2) {
			return false
		}
		return slices.Max(sl1) == slices.Max(sl2)
	})
	fmt.Println(res)
	return res
}
