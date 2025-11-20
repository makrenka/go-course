package lessons

import "fmt"

func PlayWithSlice(slice []int) []int {
	res := make([]int, len(slice))
	copy(res, slice)

	for i := len(res) - 1; i >= 0; i-- {
		if res[i] > 10 {
			befor := res[:i+1]
			after := append([]int{100}, res[i+1:]...)
			res = append(befor, after...)
			break
		}
	}

	sum := 0
	for _, v := range res {
		sum += v
	}
	if sum > 100 {
		res = append(res, 500)
	}

	evensCount := 0
	oddsCount := 0
	for _, v := range slice {
		if v%2 == 0 {
			evensCount++
		} else {
			oddsCount++
		}
	}
	if evensCount > oddsCount {
		res = append([]int{1000}, res...)
	}

	fmt.Print(res)
	return res
}
