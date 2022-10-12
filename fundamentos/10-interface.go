package main

import "fmt"

// interface
type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return (c.raio * c.raio) * 3.14
}

func (c retangulo) area() float64 {
	return (c.altura * c.largura)
}

// Função que recebe interfaces
func escreverArea(self forma) {
	fmt.Println("Area da formula é %0.2f", self.area())
}

func escreverAreaGenerica(self interface{}) {
	fmt.Println("Area da formula é %0.2f", self)
}

func main() {
	r := retangulo{10, 54}
	g := circulo{27}
	escreverArea(r)
	escreverArea(g)
	escreverAreaGenerica(g)

}
