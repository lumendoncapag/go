package main

import (
	"fmt"
)

func main() {
	//Numeros inteiros
	// int8 , int16  , int32 , int64 => São diferenciais pelo seu tamanho de bit
	// se voce declarar int sem especificar o tipo , ele vai basear na arquitetura do seu hardware.
	var numero int16 = 100
	var numero2 int = 1656
	fmt.Println(numero)
	fmt.Println(numero2)
	// uint - Numero nao tem sinal
	var numero3 int = -100
	var numero4 uint = 100
	fmt.Println(numero3)
	fmt.Println(numero4)

	//ALIAS
	// INT32 = RUNE
	// UINT8 = BYTE
	var n8 rune = 100
	var n9 byte = 222
	println(n8, n9)

	//REAIS
	//float32 e float64
	var (
		n10 float32 = 2.22
		n11 float64 = 2.222222222222222
	)
	fmt.Println(n10, n11)

	
}
