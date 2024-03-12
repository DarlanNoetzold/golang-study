package main

import (
	"fmt"
)

func main() {
	fmt.Println(retornaumint())
	fmt.Println(retornauminteumastring())
}

func retornaumint() int {
	return 10
}

func retornauminteumastring() (int, string) {
	return 20, "vinte"
}
