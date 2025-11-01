package lessons

import (
	"fmt"
	"math"
)

func PrintNumberInfo(num int) {
	if math.Abs(float64(num)) != float64(num) {
		fmt.Printf("Число %d отрицательное.\n", num)
	} else if num == 0 {
		fmt.Printf("Число равно %d.\n", num)
	} else {
		fmt.Printf("Число %d положительное.\n", num)

	}

	if num%2 != 0 {
		fmt.Printf("Число %d нечетное.\n", num)
	} else {
		fmt.Printf("Число %d четное.\n", num)
	}

	sqrt := math.Sqrt(float64(num))
	if math.Abs(float64(num)) == float64(num) && num%2 != 0 {
		fmt.Printf("Квадратный корень числа %d не является целым числом и равен %.5f.\n", num, sqrt)
	} else if math.Abs(float64(num)) == float64(num) && num%2 == 0 && num != 0 {
		fmt.Printf("Квадратный корень числа %d является целым числом и равен %d.\n", num, int(sqrt))
	}
}
