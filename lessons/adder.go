package lessons

import "fmt"

func Adder(n int) (func(x int), error) {
	res := n
	fmt.Println(res)
	return func(x int) {
		res = n + x
	}, nil
}
