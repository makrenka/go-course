package lessons

import (
	"fmt"
	"slices"
	"strings"
)

func CountVotes(votes map[string]int) string {
	if len(votes) == 0 {
		return "Кандидаты потерялись."
	}

	votesCount := make([]int, 0)
	invertedMap := make(map[int][]string)
	for key, v := range votes {
		invertedMap[v] = append(invertedMap[v], key)
		votesCount = append(votesCount, v)
	}
	maxVotes := slices.Max(votesCount)

	if maxVotes == 0 {
		return "Все голоса похищены!"
	}

	candidatesWithMaxVotes := invertedMap[maxVotes]
	slices.Sort(candidatesWithMaxVotes)

	fmt.Println(strings.Join(candidatesWithMaxVotes, ", "))
	return strings.Join(candidatesWithMaxVotes, ", ")
}
