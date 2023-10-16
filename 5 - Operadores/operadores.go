package main

import "fmt"

func main() {
	//ARITIMETRICOS
	// + , - , / , * , %

	soma := 1 + 2
	sub := 1 - 2
	div := 1 / 2
	multi := 1 * 2
	resto := 10 % 2

	fmt.Println(soma, sub, div, multi, resto)

	//Observação : Não pode fazer qualquer opreção em variaveis com tipos diferentes
	var numero1 int8 = 1
	var numero2 int32 = 1
	//ERRO
	//soma := numero1 + numero2
	//Correto
	soma2 := numero1 + int8(numero2)
	fmt.Println(soma2)

	// ATRIBUIÇÂO
	// = ou :=

	//RELACIONAIS
	// < , > , >= , <= , != , ==
	fmt.Println(1 > 2, 1 < 2, 1 != 2, 2 == 2)

	//OPERADORES LOGICOS
	// && ,  ||  , ! (not)
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso, verdadeiro || falso, !verdadeiro)

	//OPERADORES UNARIOS
	numero := 10
	numero++
	numero += 15
	numero--
	numero -= 15
	fmt.Println(numero)
}
