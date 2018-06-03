package main

import (
	"fmt"
	"time"
)

func print(from string, value int) {
	fmt.Println(from, "=>", value)
}

func main() {

	fmt.Println("---- START ----")

	for i := 0; i < 10; i++ {
		print("SEC", i)
	}

	for i := 0; i < 10; i++ {
		go print("GOR", i)
	}

	go func(msg string) {
		fmt.Println(msg)
	}("E")

	time.Sleep(2000 * time.Millisecond)
	fmt.Println("---- END ----")
}
