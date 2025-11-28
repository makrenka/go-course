package practices

import (
	"fmt"
	"maps"
	"slices"
)

func PrintRecommendations(m map[string]map[string]float64) {
	filteredMap := make(map[string]map[string]float64, len(m))

	for keyM, inner := range m {
		newInner := make(map[string]float64)
		for keyInner, val := range inner {
			newInner[keyInner] = val
		}
		filteredMap[keyM] = newInner
	}

	genres := make([]string, 0)
	for key, genre := range filteredMap {
		maps.DeleteFunc(genre, func(key string, val float64) bool {
			return val < 7
		})
		genres = append(genres, key)
	}
	slices.Sort(genres)

	for _, genre := range genres {
		fmt.Printf("%s: %s\n", genre, printMovies(filteredMap[genre]))
	}
}

func printMovies(m map[string]float64) string {
	movies := make([]string, 0)
	moviesMap := make(map[string]float64)

	for key, val := range m {
		movies = append(movies, key)
		moviesMap[key] = val
	}

	slices.SortFunc(movies, func(a, b string) int {
		if moviesMap[a] > moviesMap[b] {
			return -1
		}
		if moviesMap[a] < moviesMap[b] {
			return 1
		}
		if moviesMap[a] == moviesMap[b] {
			if a < b {
				return -1
			}
			if a > b {
				return 1
			}
			return 0
		}
		return 0
	})

	var res string
	for i := 0; i < len(movies); i++ {
		if i < len(movies)-1 {
			res = fmt.Sprintf("%s (%.1f), ", movies[i], moviesMap[movies[i]])
		} else {
			res += fmt.Sprintf("%s (%.1f).", movies[i], moviesMap[movies[i]])
		}

	}

	return res
}
