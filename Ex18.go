package main

import (
	"fmt"
)

func main() {

	anoemqueeunasci := 1988
	anoateoqualeuquerocontar := 2088

	for {
		if anoemqueeunasci > anoateoqualeuquerocontar {
			break
		}
		fmt.Println(anoemqueeunasci)
		anoemqueeunasci++
	}
}
