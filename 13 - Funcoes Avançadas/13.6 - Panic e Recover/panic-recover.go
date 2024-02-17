package main

func recuperarPanic() {
	//Com esse recover , ele nao deixa a execução do scprit morrer quando acontece um panic. Segue o jogo!
	//Quando o recover é chamado sem acontecer o panic , nao tem problemas , ele nao faz nada.
	if r := recover(); r != nil {
		println("Função recuperada com sucesso")
	}
}

func media(n1, n2 float32) bool {
	//Setando o recuperar como defer , para executar sempre no final dessa funcao.
	defer recuperarPanic()
	println("Calculando a media entre as notas .....")

	media := (n1 + n2) / 2

	//Nesse formato , quando o valor da media é 6 o script nao sabe mais oque fazer. Entra em "PANIC"
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	//Aqui estabelecemos uma msg quando acontece o panic ,atraves dessa funcao
	//Ao entrar em panic , ele chama todos os defer
	panic("A MEDIA É EXATAMENTE 6!")
}

func main() {
	println(media(6, 7))
	println("Pós execução")
}
