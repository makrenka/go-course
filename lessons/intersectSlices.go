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
		for j, v2 := range slice2 {
			if slice2[j] != len(slice2)-1 && slice2[j] == slice2[j+1] {
				continue
			}
			if v2 == v1 {
				res1 = append(res1, v2)
			}
		}
	}

	var res2 []int
	for _, v1 := range slice2 {
		for j, v2 := range slice1 {
			if slice1[j] != len(slice1)-1 && slice1[j] == slice1[j+1] {
				continue
			}
			if v2 == v1 {
				res2 = append(res2, v2)
			}
		}
	}

	fmt.Print(res1, res2)
	return res2, nil
}
