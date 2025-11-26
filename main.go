package main

import (
	"fmt"
	"go-course/practices"
)

func main() {
	m := map[string]map[string]float64{
		"Экшен": {
			"Фильм1": 8.52,
			"Фильм2": 6.0,
		},
		"Драма": {
			"Фильм3": 7.524,
			"Фильм4": 7.527,
			"Фильм5": 5.54,
		},
	}
	practices.PrintRecommendations(m)
}

func PrintRecommendations(m map[string]map[string]float64) {
	// genres := make([]string, 0)
	// movies := make([]string, 0)
	moviesMap := make(map[string]float64)
	genresMap := make(map[string]string)

	for key, val := range m {
		for k, v := range val {
			if v >= 7 {
				// genres = append(genres, key)
				// movies = append(movies, k)
				moviesMap[k] = v
				genresMap[key] = k
			}
		}
	}

	// slices.Sort(genres)
	// genres = slices.Compact(genres)
	// slices.SortFunc(movies, func(a, b string) int {
	// 	if moviesMap[a] < moviesMap[b] {
	// 		return -1
	// 	}
	// 	if moviesMap[a] > moviesMap[b] {
	// 		return 1
	// 	}
	// 	return 0
	// })

	// for _, v := range genres {
	// 	fmt.Printf("%s: %s", v, strings.Join(movies, ", "))
	// }
	fmt.Println(genresMap)
}
