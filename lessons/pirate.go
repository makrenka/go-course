package lessons

import "fmt"

var trapsNumber int
var stovesNumber int
var isDead bool

func MovePirate(isTrap bool) {
	if isDead {
		return
	}

	stovesNumber += 1
	fmt.Printf("Пират переместился на плиту %d\n", stovesNumber)

	if isTrap {
		trapsNumber += 1
	}

	if isTrap && trapsNumber <= 2 {
		fmt.Println("Пират ранен")
	} else if isTrap && trapsNumber > 2 {
		fmt.Println("Пират убит")
		isDead = true
		return
	}

	if stovesNumber == 10 {
		fmt.Println("Пират преодолел все ловушки")
		return
	}
}
