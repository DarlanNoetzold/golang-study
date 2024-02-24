package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	pessoa1 := pessoa{
		nome:      "Renata",
		sobrenome: "Pimentão",
		sabores:   []string{"pistache", "morango", "baunilha"}}

	pessoa2 := pessoa{"Frederico", "da Prússia",
		[]string{"sabão em pó", "pé de coelho", "feijão"}}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.sabores {
		fmt.Println("\t", v)
	}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, v := range pessoa2.sabores {
		fmt.Println("\t", v)
	}

}