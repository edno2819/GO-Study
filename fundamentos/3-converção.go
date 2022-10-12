package main

import "fmt"

var car string

const (
	carro   = "porsh"
	modelo  = "testa"
	tamanho = 45
	test    = iota
)

func main() {
	var name string = "casa"
	const PATH = "REDIS://google.cloud.com/banquinho/admin"
	sb := []byte(name)

	fmt.Println(sb)

	for i := 0; i < len(name); i++ {
		fmt.Printf("%v- %T - %#U - %#x\n", name[i], name[i], name[i], name[i])
	}

}
