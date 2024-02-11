package main

import (
	"fmt"
)

func main() {
	for x := 10; x <= 100; x++ {
		fmt.Println(x % 4)
	}
}