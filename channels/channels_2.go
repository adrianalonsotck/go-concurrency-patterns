package main

import (
	"time"
)

func sendString(ch chan<- string, s string) {
	ch <- s
}

func receiver(helloCh, goodbyeCh <-chan string, exitChannel chan<- bool) {
	for {
		select {
		case msg := <-helloCh:
			println(msg)
		case msg := <-goodbyeCh:
			println(msg)
		case <-time.After(time.Second * 2):
			println("Nothing received in 2 seconds. Exiting")
			exitChannel <- true
			break
		}
	}
}

func main() {
	helloCh := make(chan string, 1)
	goodbyeCh := make(chan string, 1)
	exitChannel := make(chan bool)
	go receiver(helloCh, goodbyeCh, exitChannel)
	go sendString(helloCh, "Hello The Cocktail!")
	time.Sleep(time.Second)
	go sendString(goodbyeCh, "Bye The Cocktail!")
	<-exitChannel
}
