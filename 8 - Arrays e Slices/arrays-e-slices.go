package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("---Arrays---")

	//Criando um array e depois setando um valor na posição
	var array1 [5]string
	fmt.Println(array1)
	array1[0] = "posição 1"
	fmt.Println(array1)

	//Criando um array e ja setando os valores
	array2 := [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(array2)

	//Criando um array deixando o tamanho implicito
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	fmt.Println("---Slices---")

	//Criando slice . é uma lista que não tem limitação de tamanho na hora da criação
	slice := []int{10, 1, 2, 3}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	//Append no slice
	slice = append(slice, 4)
	fmt.Println(slice)

	slice2 := array2[1:3] //valor 1 até 2 , nao inclui o 3
	fmt.Println(slice2)

	//Slice é um ponteiro para um Array , podemos então alterar o valor do array que o valor irá ser replicado no slice
	array2[1] = "3333"
	fmt.Println(slice2)

	//ARRAYS INTERNOS
	fmt.Println("---Arrays Internos---")
	slice3 := make([]int, 10, 10)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //Tamanho
	fmt.Println(cap(slice3)) //Capacidade
	slice3 = append(slice3, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //Tamanho
	fmt.Println(cap(slice3)) //Capacidade
}
