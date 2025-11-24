package lessons

import "fmt"

func CountMaxFrequency(sl []int) int {
	m := make(map[int]int)
	for _, v := range sl {
		m[v]++
	}
	maxCountKey := 0
	for _, value := range m {
		if value > maxCountKey {
			maxCountKey = value
		}
	}
	fmt.Println(maxCountKey)
	return maxCountKey
}
