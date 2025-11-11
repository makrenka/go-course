package lessons

import "fmt"

func PrintReplaced(str string) {
	for _, r := range str {
		sym := string(r)
		if sym == "у" {
			sym = "а"
		}
		fmt.Print(sym)
	}
}
