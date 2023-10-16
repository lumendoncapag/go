package main

import "fmt"

// Funcão soma
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Nesse metodo de declaração todos as variaveis que nao fornecer o tipo , irá obedecer o da direita
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 30)
	fmt.Println(soma)

	//Aninhar funcções em variaveis
	var f = func(txt string) {
		fmt.Println("Printando " + txt)
	}
	f("tudo")

	//Chamando os dois resultados
	resultadoSum, resultadoSub := calculosMatematicos(10, 15)
	fmt.Println(resultadoSum, resultadoSub)

	//Chamando apenas um dos resultados da funcao
	resultadoSoma, _ := calculosMatematicos(10, 15)
	_, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

}
