package main

import (
	"fmt"
	"time"
)

func write(texto string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value received %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}

// O padrão generator é muito utilizado quando se quer esconder a complexidade da go rotine
func main() {
	channel := write("Hellou world")

	for i := 0; i < 20; i++ {
		fmt.Println(<-channel)
	}
}
