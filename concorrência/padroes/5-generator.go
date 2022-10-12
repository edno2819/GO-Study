package main

import (
	"fmt"
	"time"
)

// O padrão generator é muito utilizado quando se quer esconder a complexidade da go rotine
func main() {
	canal := escrever("OLá mundo")

	for i := 0; i < 20; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
