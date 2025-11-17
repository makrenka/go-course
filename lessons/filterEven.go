package lessons

import "fmt"

func FilterEven(nums ...int) (res []int) {
	for i := range nums {
		if nums[i]%2 == 0 {
			res = append(res, nums[i])
		}
	}
	fmt.Print(res)
	return
}
