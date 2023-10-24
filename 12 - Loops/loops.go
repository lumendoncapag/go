package main

import "time"

func main() {
	println("---Loops---")

	//Loop for com declaração de variavel "global" , ela por
	i := 0
	for i < 1 {
		i++
		println("incremente i")
		time.Sleep(time.Second)
	}
	println(i)

	//Loop for com declaração init
	for j := 0; j < 1; j++ {
		println("Incrementando J", j)
		time.Sleep(time.Second)
	}

	//Iterando uma Array com range
	nomes := [3]string{"Joao", "Davi", "Lucas"}
	for indice, nome := range nomes {
		println(indice, nome)
	}

	//Range em apenas um dos valores
	for _, nome := range nomes {
		println(nome)
	}

	slice_nomes := nomes
	for indice := range slice_nomes {
		println(indice)
	}

	//Iterando uma string
	for indice, letra := range "PALAVRA" {
		println(indice, string(letra))
	}

	println("---Iterando MAP ---")

	usuario := map[string]string{
		"nome":      "Lucas",
		"sobrenome": "Alberto",
	}

	for key, value := range usuario {
		println(key, value)
	}

	//Loop infinito , mais conhecido com While true
	for {
		time.Sleep(2)
	}
}
