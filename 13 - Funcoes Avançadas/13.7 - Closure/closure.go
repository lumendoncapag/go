package main

func closure() func() {
	texto := "Dentro da funcao Closure"

	funcaoClosure := func() {
		//Pega o valor texto dentro dessa funcao .
		println(texto)
	}
	return funcaoClosure
}

func main() {
	texto := "Dentro da funcao Main"
	println(texto)

	//Ele vai rodar a funcao que esta dentro da Closure
	funcaoMain := closure()
	funcaoMain()
}
