package main

//nomeia o valor de saida da funcao ja na declarao
func calculoMatematicos(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	//Nao precisa retornar a variavel return
	return
}

func main() {
	soma, subtracao := calculoMatematicos(2, 4)
	println(soma, subtracao)
}
