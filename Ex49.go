package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {

	zezinho := pessoa{"Zezinho", "Primodoluizinho", 10}
	fmt.Println(zezinho)
	mudeMe(&zezinho)
	fmt.Println(zezinho)
	
}

func mudeMe(p *pessoa) {
	(*p).nome = "Luizinho"
	p.sobrenome = "Primodozezinho"
}
