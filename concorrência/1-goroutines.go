package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)
	go func() {
		escrever("Olá mundo")
		waitGroup.Done()
	}()
	go func() {
		escrever("Olá mundo - 2")
		waitGroup.Done()
	}()

	waitGroup.Wait()
	fmt.Println("Finish")

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(1)
	}
}
