package practices

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func PrintRecommendations(m map[string]map[string]float64) {
	genres := make([]string, 0)
	movies := make([]string, 0)
	moviesMap := make(map[string]float64)
	genresMap := make(map[string][]string)

	for key, val := range m {
		for k, v := range val {
			if v >= 7 {
				genres = append(genres, key)
				movies = append(movies, k)
				moviesMap[k] = v
				genresMap[key] = append(genresMap[key], k)
			}
		}
	}

	slices.Sort(genres)
	genres = slices.Compact(genres) // прыбіраем жанры, што дублююцца
	slices.SortFunc(movies, func(a, b string) int {
		if moviesMap[a] < moviesMap[b] {
			return -1
		}
		if moviesMap[a] > moviesMap[b] {
			return 1
		}
		return 0
	})

	for _, v := range genres {
		fmt.Printf("%s: %s\n", v, printMovie(moviesMap))
	}
}

func printMovie(moviesMap map[string]float64) string {
	moviesWithRates := make([]string, 0)
	for key, val := range moviesMap {
		moviesWithRates = append(moviesWithRates, key+" ("+strconv.FormatFloat(val, 'f', 1, 64)+")")
	}
	return strings.Join(moviesWithRates, ", ")
}
