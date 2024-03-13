package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

func (q quadrado) area() {
	resultado := q.lado * q.lado
	fmt.Println("Área do quadrado:", resultado)
}

type circulo struct {
	raio float64
}

func (c circulo) area() {
	resultado := math.Pi * 2 * c.raio
	fmt.Println("Área do círculo:", resultado)
}

type info interface {
	area()
}

func medida(i info) {
	i.area()
}

func main() {

	x := quadrado{
		lado: 15.0,
	}

	y := circulo{
		raio: 5.25,
	}

	//x.area()
	//y.area()

	medida(x)
	medida(y)
}