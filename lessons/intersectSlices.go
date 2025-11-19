package lessons

import (
	"errors"
	"fmt"
)

// маё
func IntersectSlices(slice1, slice2 []int) ([]int, error) {
	if slice1 == nil || slice2 == nil {
		err := errors.New("slices cannot be nil")
		return nil, err
	}

	var result []int

	copySlice1 := make([]int, len(slice1))
	copy(copySlice1, slice1)
	copySlice2 := make([]int, len(slice2))
	copy(copySlice2, slice2)

	for i := 0; i < len(copySlice1); i++ {
		for j := 0; j < len(copySlice2); j++ {
			if copySlice2[j] == copySlice1[i] {
				result = append(result, copySlice2[j])
				if copySlice1[i] != len(copySlice1)-1 || copySlice2[j] != len(copySlice2)-1 {
					copySlice2 = copySlice2[j+1:]
				}
				break
			}
		}
	}

	fmt.Print(result)
	return result, nil
}

// харошае рашэньне!!!
func intersectSlices(nums1, nums2 []int) ([]int, error) {
	intersectNums := []int{}
	if nums1 == nil || nums2 == nil {
		return intersectNums, errors.New("slices cannot be nil")
	}
	idx1 := 0
	idx2 := 0
	for idx1 < len(nums1) && idx2 < len(nums2) {
		switch {
		case nums1[idx1] < nums2[idx2]:
			idx1++
		case nums1[idx1] > nums2[idx2]:
			idx2++
		default:
			intersectNums = append(intersectNums, nums1[idx1])
			idx1++
			idx2++
		}
	}
	return intersectNums, nil
}
