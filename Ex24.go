package main

import (
	"fmt"
)

func main() {

	slice := []int{10, 20, 30, 40, 50, 1, 2, 3, 4, 5}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", slice)
}
