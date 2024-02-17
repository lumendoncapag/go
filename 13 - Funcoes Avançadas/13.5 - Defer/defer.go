package main

func funcao1() {
	println("Funcao 1")
}

func funcao2() {
	println("Funcao 2")
}

func media(n1, n2 float32) bool {
	//Com esse defer , ele ira executar esse print antes da função encerrar
	defer println("Media calculada ... informando ")
	println("Calculando a media entre as notas .....")

	media := (n1 + n2) / 2

	if media >= 6 {
		//Para informar essa mensagem na tela antes dos resultados , teria que repitir antes de todos os returns
		//println("Media calculada ... informando ")
		return true
	}

	return false
}

func main() {
	defer funcao1()
	// DEFER = ADIAR
	// O defer faz com que a execução so seja feita em um ultimo momento. Normalmente antes da finalização da função
	funcao2()

	if media(6, 8) == true {
		println("Passado")
	} else {
		println("Reprovado")
	}
}
