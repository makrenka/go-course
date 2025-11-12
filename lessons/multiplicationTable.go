package lessons

import "fmt"

func PrintTable(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			if j != num {
				fmt.Printf("%d x %d = %d\t", i, j, i*j)
			} else {
				fmt.Printf("%d x %d = %d\n", i, j, i*j)
			}
		}
	}
}
