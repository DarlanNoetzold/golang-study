package main

import (
	"fmt"
)

func main() {
	essavaireceberaoutrafuncao(issovaiserumargumento)
}

func issovaiserumargumento() {
	fmt.Println("Olha eu aqui!")
}

func essavaireceberaoutrafuncao(x func()) {
	fmt.Println("Prestenção:")
	x()
}
