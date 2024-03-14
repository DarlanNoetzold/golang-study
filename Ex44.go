package main

import (
	"fmt"
)

func main() {

	x := retornaumafunc()

	x()
}

func retornaumafunc() func() {
	return func() {
		fmt.Println("Olha eu aqui!")
	}
}
