package lessons

import (
	"fmt"
	"slices"
)

func DeletingFromSlice(slice []int) []int {
	res := make([]int, len(slice))
	copy(res, slice)

	var deleted []int

	if len(res) > 0 && res[len(res)-1] > 10 {
		res = res[:len(res)-1]
		deleted = append(deleted, len(res)-1)
	}

	if len(res) > 2 && cap(res) > 5 {
		befor := res[:2]
		after := res[2+1:]
		res = append(befor, after...)
		deleted = append(deleted, res[2])
	}

	if len(res) > 0 && len(deleted) == 2 {
		res = res[1:]
	}

	res = slices.Clip(res)

	fmt.Println(res, cap(res))
	return res
}
