package main

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)

}

func main() {
	println("Funções recursivas")
	// 1 1 2 3 5 8 13 Sequencia de fibonacci
	var posicao uint = 3
	println(fibonacci(posicao))

	for i := uint(0); i < 10; i++ {
		println(fibonacci(i))
	}
}
