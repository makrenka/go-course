package lessons

import "errors"

func Calculate(n1 float64, n2 float64, action string) (float64, error) {
	if action == "add" {
		return n1 + n2, nil
	}
	if action == "subtract" {
		return n1 - n2, nil
	}
	if action == "multiply" {
		return n1 * n2, nil
	}
	if n2 == 0 {
		return 0, errors.New("division by zero")
	}
	if action == "divide" {
		return n1 / n2, nil
	}

	return 0, errors.New("unknown operation")
}
