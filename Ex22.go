package main

import (
	"fmt"
)

func main() {
	esporteFavorito := "starcraft"

	switch esporteFavorito {

	case "futebol":
		fmt.Println("quer jogar futebol, vai pro brasil")
	case "starcraft":
		fmt.Println("quer jogar starcraft, vai pra coréia")
	case "espeleísmo":
		fmt.Println("quer fazer essa coisa estranha, vai pro psiquiatra")
	}

}
