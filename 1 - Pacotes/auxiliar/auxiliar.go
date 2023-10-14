package auxiliar

import (
	"fmt"
)

// Funções que vao ser importadas devem inciar com Letra Maiscusula: Exemplo abaixo. não é escrever e sim Escrever.
// Função de exemplo para Escrever
func Escrever() {
	fmt.Println("Escrevendo do arquivo do Pacote Auxiliar")
	//chamando uma função do mesmo pacote. Não é nescessario letra maiscula nem nome do pacote
	escrever2()
}
