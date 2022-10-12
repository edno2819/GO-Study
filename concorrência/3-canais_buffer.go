package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			canal2 <- "Canal 2"
		}
	}()

	// for {
	// 	msg1 := <-canal1
	// 	fmt.Println(msg1)

	// 	msg2 := <-canal2
	// 	fmt.Println(msg2)
	// }

	//Select usado para impedir o travamennto dos canais, executa a ação dos canais já prontos
	for {
		select {
		case msg1 := <-canal1:
			fmt.Println(msg1)
		case msg2 := <-canal2:
			fmt.Println(msg2)
		}
	}

}
