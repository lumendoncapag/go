package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

//Para calcular a area, n é nescessario chamar 2 funcoes diferentes
// Pode ser criar uma assinatura para unificar o calculo da area usando interfaces

type forma interface {
	area() float64
	//Nesse caso , cada struct tera que ter um metodo com calculo de sua area
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func escrevaArea(f forma) {
	fmt.Printf("A Area da forma é %0.2f\n", f.area())
}

func main() {
	r := retangulo{15, 10}
	c := circulo{3}
	escrevaArea(r)
	escrevaArea(c)
}
