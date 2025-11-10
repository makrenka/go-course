package lessons

import (
	"fmt"
	"math/rand/v2"
)

func RollDice(num int) {
	for i := 1; ; i++ {
		cub1 := rand.IntN(6) + 1
		cub2 := rand.IntN(6) + 1
		str1 := "ок"
		str2 := "ка"
		str3 := "ков"
		str4 := "ся"
		str5 := "ось"

		if num == (cub1 + cub2) {
			if i%10 == 1 && i != 11 {
				fmt.Printf("Выпало %d и %d, в сумме %d, на это потребовал%s %d брос%s.\n", cub1, cub2, (cub1 + cub2), str4, i, str1)
				break
			} else if i%10 >= 2 && i%10 <= 4 && i <= 14 {
				fmt.Printf("Выпало %d и %d, в сумме %d, на это потребовал%s %d брос%s.\n", cub1, cub2, (cub1 + cub2), str5, i, str2)
				break
			} else {
				fmt.Printf("Выпало %d и %d, в сумме %d, на это потребовал%s %d брос%s.\n", cub1, cub2, (cub1 + cub2), str5, i, str3)
				break
			}
		} else {
			fmt.Printf("Выпало %d и %d, в сумме %d, бросаем еще раз.\n", cub1, cub2, (cub1 + cub2))
		}
	}
}
