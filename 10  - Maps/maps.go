package main

import "fmt"

func main() {
	fmt.Println("---Maps---")
	//tipo da key , tipo do valor
	usuario := map[string]string{
		"nome":      "Lucas",
		"sobrenome": "Alberto", //Sempre ter o virgula no final
	}
	//Print map inteiro
	fmt.Println(usuario)

	//Print valor especifico
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "lucas",
			"ultimo":   "ferreira",
		},
	}
	fmt.Print(usuario2)

	//Deletando uma chave
	delete(usuario, "nome")
	fmt.Println(usuario)

	//adicionar
	usuario["signo"] = "capriocornio"
	fmt.Println(usuario)
}
