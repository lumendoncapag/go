package main

import "fmt"

var n int

//Será sempre excutada primeiro , antes até da main.
//Usada para fazer configurações iniciais
func init() {
	n = 10
	fmt.Println("Função init sendo executada")
}

func main() {
	fmt.Println("Função main sendo executada")
	println(n)
}
