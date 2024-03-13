package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("isso vai aparecer depois")
	fmt.Println("isso vai aparecer antes")
}
