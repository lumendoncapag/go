package main

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinialPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	// Sem ponteiro , o numero novo será feito uma copia .
	numero := 10
	println(numero)
	numeroInvertido := inverterSinal(numero)
	println(numeroInvertido)
	println(numero)

	//Utilizando ponteiro
	novoNumero := 40
	println(novoNumero)
	inverterSinialPonteiro(&novoNumero)
	println(novoNumero)
}
