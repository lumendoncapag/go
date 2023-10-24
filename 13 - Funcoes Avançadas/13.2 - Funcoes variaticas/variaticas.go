package main

import "fmt"

// Ela vai receber de 1 a N . Os tres pontos serve para isso
func soma(numeros ...int) int {
	//Ele retorna um slice
	fmt.Println(numeros)
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		println(texto, numero)
	}
}

func main() {
	println(soma(1, 2, 3, 4, 5, 6, 7, 8, 9))
	escrever("ola mundo", 1, 2, 3, 4, 5, 6, 7)

}
