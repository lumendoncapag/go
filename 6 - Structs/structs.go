package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	//Chamando um struct dentro de outro Struct
	user_endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

//Herança - Só que não
//Um estudante é uma pessoa

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	//Não precisa declarar o tipo , igual o tipo endereçõ do exemplo acima
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("--- Arquivo Struct -----")

	//Formas de chamar um struc
	// 1 -- declarando variavel
	var u usuario
	u.nome = "Lucas"
	u.idade = 23
	fmt.Println(u)

	// 2 -- Por inferencia
	endereco2 := endereco{"Rua Eilzo", 153}
	u2 := usuario{"Lucas", 23, endereco2}
	fmt.Println(u2)

	// 3 -- Nao tem todos os dados
	u3 := usuario{nome: "Lucas"}
	fmt.Println(u3)

	//Herança - Só que não
	fmt.Println("----Herança----")
	p1 := pessoa{"Joao", "Pedro", 20, 1.65}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia Eletrica", "UFPB"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
}
