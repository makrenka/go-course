package lessons

import "fmt"

func ReplaceEvenOnEvenIndices(slice [][]int) [][]int {
	var newSlice [][]int

	for j, nestedSlice := range slice {
		for i, v := range nestedSlice {
			if i % 2 == 0 && v % 2 == 0 {
				nestedSlice[i] = 0
			} 
		}
		newSlice = append(newSlice, slice[j])
	}

	fmt.Print(newSlice)
	return newSlice
}
