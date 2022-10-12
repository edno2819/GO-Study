package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// u - variavel para reeferenciar outro campos do usuário, como se fosse um self
func (u usuario) salvar() {
	fmt.Println("Usuário salvo", u.nome)
}

// Usou o ponteiro para referenciar o local da memoria do objeto
func (self *usuario) aniversario() {
	self.idade++
}

func main() {
	usuario1 := usuario{"Edno almeida", 20}
	fmt.Println(usuario1)
	fmt.Println(usuario1.idade)
	usuario1.salvar()
	usuario1.aniversario()
	fmt.Println(usuario1.idade)

}
