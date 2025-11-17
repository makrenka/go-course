package lessons

func SumNeighbors(arr [10]int) [10]int {
	res := [10]int{}

	for i := range arr {
		if i == 0 {
			res[i] = arr[i+1]
		} else if i == len(arr)-1 {
			res[i] = arr[i-1]
		} else {
			res[i] = arr[i-1] + arr[i+1]
		}
	}

	return res
}
