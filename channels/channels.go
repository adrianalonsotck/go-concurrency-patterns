package main

import "fmt"

func main() {
	channel := make(chan string)
	go func() {
		channel <- "Hello The Cocktail!"
	}()

	fmt.Println("Waiting channel")
	message := <-channel
	fmt.Println(message)
}
