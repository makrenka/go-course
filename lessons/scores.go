package lessons

import (
	"errors"
	"fmt"
)

func GetLetterScore(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func GetScore() (int, error) {
	var score int
	fmt.Printf("Увядзіце вашу адзнаку ад 0 да 100: ")

	if _, err := fmt.Scan(&score); err != nil {
		return 0, errors.New("памылка ўвода: увядзіце цэлы лік")
	}

	if score < 0 || score > 100 {
		return 0, errors.New("адзнака павінна быць ад 0 да 100")
	}

	return score, nil
}
