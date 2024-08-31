package main

import (
	"fmt"
	"time"
)

// The channel only receive a new value whether the channel is void.
func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text // Sending a value to into channel
		fmt.Println("Adding msg in channel!")
	}
	close(channel)
}

func main() {
	channel := make(chan string)
	go write("Receiving - Hellou world", channel)

	for {
		menssage, open := <-channel // Waiting the channel to receive a value
		if !open {
			break
		}
		fmt.Println(menssage)
		time.Sleep(time.Second * 1)
	}
	fmt.Println("After channel!")
}
