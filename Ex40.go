package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	sobrenome  string
	idade int
}

func (ODEMOGORGONDATV pessoa) demonstrar() {
	fmt.Println("O nome completo dessa pessoa é:", ODEMOGORGONDATV.nome, ODEMOGORGONDATV.sobrenome, "\nEssa pessoa tem", ODEMOGORGONDATV.idade, "anos.")
}

func main() {

	aguriazinhadecabeloraspadodatv := pessoa {
		nome: "Doze",
		sobrenome: "Esquisita",
		idade: 13,
	}
	
	aguriazinhadecabeloraspadodatv.demonstrar()	

}
