package lessons

import "fmt"

func PrintDiamond(n int) {
	fmt.Println("Мой бриллиант:")
	space := ""

	for i := range n {
		spaceCount := n - i - 1
		if i == 0 {
			spaces := fmt.Sprintf("%*s", spaceCount, space)
			fmt.Printf("%s#\n", spaces)
		} else {
			firstSpaces := fmt.Sprintf("%*s", spaceCount, space)
			secondSpaceCount := 2*i - 1
			secondSpaces := fmt.Sprintf("%*s", secondSpaceCount, space)
			fmt.Printf("%s#%s#\n", firstSpaces, secondSpaces)
		}
	}

	for i := n - 2; i >= 0; i-- {
		spaceCount := n - i - 1
		if i == 0 {
			spaces := fmt.Sprintf("%*s", spaceCount, space)
			fmt.Printf("%s#\n", spaces)
		} else {
			firstSpaces := fmt.Sprintf("%*s", spaceCount, space)
			secondSpaceCount := 2*i - 1
			secondSpaces := fmt.Sprintf("%*s", secondSpaceCount, space)
			fmt.Printf("%s#%s#\n", firstSpaces, secondSpaces)
		}
	}
}
