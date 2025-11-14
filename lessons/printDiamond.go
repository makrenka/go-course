package lessons

import "fmt"

// func PrintDiamond(n int) {
// 	fmt.Println("Мой бриллиант:")
// 	space := ""

// 	for i := range n {
// 		spaceCount := n - i - 1
// 		if i == 0 {
// 			spaces := fmt.Sprintf("%*s", spaceCount, space)
// 			fmt.Printf("%s#\n", spaces)
// 		} else {
// 			firstSpaces := fmt.Sprintf("%*s", spaceCount, space)
// 			secondSpaceCount := 2*i - 1
// 			secondSpaces := fmt.Sprintf("%*s", secondSpaceCount, space)
// 			fmt.Printf("%s#%s#\n", firstSpaces, secondSpaces)
// 		}
// 	}

// 	for i := n - 2; i >= 0; i-- {
// 		spaceCount := n - i - 1
// 		if i == 0 {
// 			spaces := fmt.Sprintf("%*s", spaceCount, space)
// 			fmt.Printf("%s#\n", spaces)
// 		} else {
// 			firstSpaces := fmt.Sprintf("%*s", spaceCount, space)
// 			secondSpaceCount := 2*i - 1
// 			secondSpaces := fmt.Sprintf("%*s", secondSpaceCount, space)
// 			fmt.Printf("%s#%s#\n", firstSpaces, secondSpaces)
// 		}
// 	}
// }

//---------------------------------------------
// Харошы варыянт

// func PrintDiamond(n int) {
// 	fmt.Println("Мой бриллиант:")
// 	n = n - 1
// 	left := n
// 	right := n
// 	for i := 0; i <= n*2; i++ {
// 		for j := 0; j <= n*2; j++ {
// 			if j == left || j == right {
// 				fmt.Print("#")
// 			} else if j < right {
// 				fmt.Print(" ")
// 			}
// 		}
// 		if i < n {
// 			left--
// 			right++
// 		} else {
// 			left++
// 			right--
// 		}

// 		fmt.Println()
// 	}
// }

//-------------------------------------------------
// ChatGPT, таксама харошы, з укладзеным цыклам

func PrintDiamond(n int) {
	fmt.Println("Мой бриллиант:")

	// Верхняя частка ромба
	for i := 0; i < n; i++ {
		// Перад першым # трэба надрукаваць (n - i - 1) прабелаў
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}

		// Друкуем левы бок #
		fmt.Print("#")

		// Калі i > 0, то паміж дзвюма # ёсць унутраныя прабелы (2*i - 1)
		if i > 0 {
			for j := 0; j < 2*i-1; j++ {
				fmt.Print(" ")
			}
			fmt.Print("#")
		}

		// Канец радка
		fmt.Println()
	}

	// Ніжняя частка ромба
	for i := n - 2; i >= 0; i-- {
		// Перад першым #
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}

		// Друкуем левы #
		fmt.Print("#")

		// Унутраныя прабелы (толькі калі не сярэдзіна)
		if i > 0 {
			for j := 0; j < 2*i-1; j++ {
				fmt.Print(" ")
			}
			fmt.Print("#")
		}

		fmt.Println()
	}
}
