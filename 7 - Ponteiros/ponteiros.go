package main

import "fmt"

func main() {
	fmt.Println("----Ponteiros----")
	//Nesse formato , a variavel2 vai receber uma copia da variavel 1
	var variavel int = 10
	var variavel2 int = variavel
	fmt.Println(variavel2)
	//Quando houve qqualquer modificação na variavel1 nao irá ser refletida na variavel2, pois ja foi feito a copia anterior
	variavel++
	fmt.Println(variavel, variavel2)

	//Ponteiro é a Referencia da Memoria daquela variavel

	//aplicação ponteiro
	var variavel3 int = 10
	var ponteiro *int = &variavel3
	fmt.Println(ponteiro)
	variavel3++
	//Apontando para o endereçõ de memoria
	fmt.Println(variavel3, ponteiro)

	//Apontando para o valor no endereço de memoria
	fmt.Println(variavel3, *ponteiro)

}
