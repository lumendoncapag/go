package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// Criando metodo para o struct usuario, exemplo de salvando nome no banco
// u = Variavel que vai se referenciar ao struct para obter seus dados, exemplo. u.idade , u.nome
// salvar() = nome do Metodo

//METODO
func (u usuario) salvar() {
	fmt.Printf("Salvando o User: %s no banco\n", u.nome)
}

//bool = tipo de retorno do metodo. Mesmo formato utilizado em funcao
func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

//Metodo para alterar valor de um atributo de um struct
//Utilizar ponteiro para alterar o valor raiz e nao uma copia
func (u *usuario) aniversario() {
	u.idade++
}
func main() {
	usuario1 := usuario{"Usuario 1", 20}
	fmt.Println(usuario1)
	//Chamando O Metodo salvar() do struc usuario
	usuario1.salvar()

	usuario2 := usuario{"Lucas", 17}
	usuario2.salvar()

	println(usuario2.maiorDeIdade()) //17

	usuario2.aniversario() //18

	println(usuario2.maiorDeIdade())
}
