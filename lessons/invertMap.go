package lessons

import (
	"fmt"
	"slices"
)

func InvertMap(m map[string]int) map[int][]string {
	res := make(map[int][]string)
	for key, value := range m {
		res[value] = append(res[value], key)
	}
	return res
}

func PrintMap(m map[int][]string) {
	keys := make([]int, 0, len(m))
	for k, v := range m {
		slices.Sort(v)
		keys = append(keys, k)
	}

	slices.Sort(keys)

	fmt.Println("{")
	for _, k := range keys {
		fmt.Printf("  %d: [", k)
		values := m[k]
		for i, v := range values {
			fmt.Printf("\"%s\"", v)
			if i < len(values)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Println("],")
	}
	fmt.Println("}")
}
