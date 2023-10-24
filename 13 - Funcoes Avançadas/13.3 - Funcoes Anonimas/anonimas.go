package main

import "fmt"

func main() {
	//Funcao anonima
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Ola mundo")

	println(retorno)

}
