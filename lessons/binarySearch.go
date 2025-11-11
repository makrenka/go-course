package lessons

import "errors"

var guesses int
var random int

func guess(num int) (int, error) {
	if guesses >= 6 {
		return 0, errors.New("too many attempts")
	}
	guesses++
	if num > random {
		return -1, nil
	}
	if num < random {
		return 1, nil
	}
	return 0, nil
}

func Play() int {
	low, high := 1, 100

	for low <= high {
		mid := (low + high) / 2
		res, _ := guess(mid)

		if res == 0 {
			return mid
		}

		if res == 1 {
			low = mid + 1
		}

		if res == -1 {
			high = mid - 1
		}
	}

	return low
}
