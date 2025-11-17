package lessons

import (
	"errors"
	"fmt"
)

func Max(nums []int) (int, error) {
	if nums == nil {
		err := errors.New("slice is nil or empty")
		fmt.Println("error:", err)
		return 0, err
	}

	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Print(max)
	return max, nil
}
