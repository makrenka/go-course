package lessons

import (
	"errors"
	"fmt"
)

func IntersectSlices(slice1, slice2 []int) ([]int, error) {
	if slice1 == nil || slice2 == nil {
		err := errors.New("slices cannot be nil")
		return nil, err
	}

	var res1 []int
	for _, v1 := range slice1 {
		for _, v2 := range slice2 {
			if v2 == v1 {
				res1 = append(res1, v2)
			}
		}
	}

	var res2 []int
	for _, v1 := range slice2 {
		for _, v2 := range slice1 {
			if v2 == v1 {
				res2 = append(res2, v2)
			}
		}
	}

	fmt.Print(res1, res2)
	return res2, nil
}
