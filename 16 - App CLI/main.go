package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()
	// erro := aplicacao.Run(os.Args)
	// if erro != nil {
	// 	log.Fatal(erro)
	// }

	//Jeito mais pratico de delcarar o if
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
