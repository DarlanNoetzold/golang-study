package main

import (
	"fmt"
)

type veículo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veículo
	traçãoNasQuatro bool
}

type sedan struct {
	veículo
	modeloLuxo bool
}

func main() {
	carrãodotio := sedan{veículo{4, "abóbora"}, true}
	fubicadovô := caminhonete{
		veículo: veículo{
			portas: 8,
			cor:    "ferrugem",
		},
		traçãoNasQuatro: false,
	}

	fmt.Println(carrãodotio)
	fmt.Println(fubicadovô)
	fmt.Println(carrãodotio.cor)
	fmt.Println(fubicadovô.traçãoNasQuatro)

}