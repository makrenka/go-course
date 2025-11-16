package lessons

import (
	"fmt"
)

func GenerateSecretCode(arr [5]int) string {
	min, max := arr[0], arr[0]
	var strWithPrefixes string

	for _, v := range arr {
		if v%2 == 0 {
			strWithPrefixes += fmt.Sprintf("E%d", v)
		} else {
			strWithPrefixes += fmt.Sprintf("O%d", v)
		}

		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}
	res := fmt.Sprintf("%d%s%d", min, strWithPrefixes, max)
	fmt.Print(res)

	return fmt.Sprintf("%d%s%d", min, strWithPrefixes, max)
}
