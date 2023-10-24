package main

import (
	"fmt"
)

func main() {
	fmt.Println("---Estruturas de controle---")

	numero := -11
	//IF-ELSE basico
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}
	//Declarando variavel durante a condição
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println(" É maior que 0")
	} else if outroNumero < -10 {
		print(outroNumero)
		fmt.Println(" É menor ou igual a -10")
	} else {
		print(outroNumero)
		fmt.Printf(" Nao é maior que 0 ou menor que -10")
	}
	//Essa variavel só fica disponivel durante o uso do escopo do if
	//fmt.Println(outroNumero)

	switchFunc()
}

// SWITCH

func diaDaSemana(numero int) string {
	var diaDaSemana string
	switch numero {
	case 1:
		diaDaSemana = "Domingo"
		fallthrough //Clausula que Joga  pra dentro da proxima condição , ou seja , se for domingo sempre vai printar Segunda
	case 2:
		diaDaSemana = "Segunda"
	case 3:
		diaDaSemana = "Terça"
	case 4:
		diaDaSemana = "Quarta"
	case 5:
		diaDaSemana = "Quinta"
	case 6:
		diaDaSemana = "Sexta"
	case 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Insira um numero de 1 a 7"
	}
	return diaDaSemana
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	default:
		return "insira um numero valido"
	}
}

func switchFunc() {
	fmt.Println("--- Switch---")
	//chamando a funcao
	fmt.Println(diaDaSemana(1))

	//com variavel
	dia := diaDaSemana(2)
	println(dia)

	dia2 := diaDaSemana2(2)
	println(dia2)

}
