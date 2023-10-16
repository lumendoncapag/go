package main

import "fmt"

func main() {

	//Explicitando tipo da variavel
	var variavel1 string = "Variavel 1"
	fmt.Println(variavel1)

	//Declarando variavel com tipo implicito
	variavel2 := "Variavel 2"
	fmt.Println(variavel2)

	//Declarar em um bloco
	var (
		variavel3 string = "Var 3"
		variavel4 string = "Var 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Var 5", "Var 6"

	fmt.Println(variavel5, variavel6)

}
