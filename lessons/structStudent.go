package lessons

import (
	"fmt"
	"math"
)

type Student struct {
	Name   string
	Grades []int
}

func (s Student) AverageGrade() float64 {
	if len(s.Grades) == 0 {
		return 0
	}

	var sum float64
	for _, v := range s.Grades {
		sum += float64(v)
	}

	return math.Round(sum/float64(len(s.Grades))*10) / 10
}

func (s Student) Info() string {
	res := fmt.Sprintf("Студент %s, средняя оценка: %.1f.", s.Name, s.AverageGrade())

	return res
}
