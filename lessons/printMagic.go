package lessons

import (
	"fmt"
	"strings"
)

func PrintMagic(slice []int) {
	res := make([]int, len(slice), 5)
	for i := range res {
		res[i] = 1
		for j := range slice {
			if j != i {
				res[i] *= slice[j]
			}
		}
	}

	fmt.Println(strings.ReplaceAll(fmt.Sprint(res), " ", ", "))
}
