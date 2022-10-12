package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("ola mundo", canal)

	for {
		mensagem, aberto := <-canal // Esperando o canla receber um valor
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}
	fmt.Println("Depois do canal")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		canal <- texto // Mandando um valor para dentro do canl
	}
	close(canal)
}
