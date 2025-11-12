package lessons

import "fmt"

func PrintDiamond(n int) {
	fmt.Println("Мой бриллиант:")
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			// space := ""
			// if i == j && (i == 1 || i == n) {
			// 	spaceCount := n - 1
			// 	spaces := fmt.Sprintf("%*s", spaceCount, space)
			// 	fmt.Printf("%s#\n", spaces)
			// } else if i < j {
			// 	spaceCount := n - i - 1
			// 	spaces := fmt.Sprintf("%*s", spaceCount, space)
			// 	fmt.Printf("%s#", spaces)
			// }
			fmt.Printf("i:%d j:%d\n", i, j)
		}
	}
}
