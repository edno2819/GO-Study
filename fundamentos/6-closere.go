package main

import "fmt"

//uma closure lhe dá acesso ao escopo de uma função externa a partir de uma função interna.

func closure(texto string) func() {
	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	texto := "Dentro da função main"
	funcaoNova := closure(texto)
	funcaoNova()
}
