package lessons

import (
	"fmt"
	"maps"
)

// маё
// func MergeMaps(m1, m2 map[string]int) map[string]int {
// 	clone1 := maps.Clone(m1)
// 	clone2 := maps.Clone(m2)
// 	for k1, v1 := range clone1 {
// 		for k2, v2 := range clone2 {
// 			if k2 == k1 {
// 				clone2[k2] = v2 + v1
// 			}
// 		}
// 	}
// 	maps.Copy(clone1, clone2)
// 	fmt.Println(clone1)
// 	return clone1
// }

// лепшае
func MergeMaps(m1, m2 map[string]int) map[string]int {
	res := maps.Clone(m1)

	for key, val := range m2 {
		res[key] += val
	}

	fmt.Println(res)
	return res
}
