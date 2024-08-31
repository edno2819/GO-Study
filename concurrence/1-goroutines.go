package main

import (
	"fmt"
	"sync"
	"time"
)

func write(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		write("Hellou world")
		waitGroup.Done()
	}()

	go func() {
		write("Hellou world - 2")
		waitGroup.Done()
	}()

	waitGroup.Wait()
	fmt.Println("Finish!")
}
