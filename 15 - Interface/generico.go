package main

import "fmt"

// É utilizada para facilitar o uso de um função sem nescessariamente
//especificar o time do que vai utilizar
//Nesse exemplo , ele irá printar qualque type: String, int , bool e etc
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
}
