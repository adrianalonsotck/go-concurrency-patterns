package main

import (
	"fmt"
)

func generator(max int) <-chan int {

	outChInt := make(chan int, 100)
	go func() {
		defer close(outChInt)

		for i := 1; i <= max; i++ {
			outChInt <- i
		}
	}()
	return outChInt
}

func power(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		defer close(out)

		for v := range in {
			out <- v * v
		}
	}()
	return out
}

func sum(in <-chan int) <-chan int {

	out := make(chan int, 100)
	go func() {
		defer close(out)

		var sum int
		for v := range in {
			sum += v
		}
		out <- sum
	}()
	return out
}

func LaunchPipeline(amount int) int {
	firstCh := generator(amount)
	secondCh := power(firstCh)
	thirdCh := sum(secondCh)
	result := <-thirdCh
	return result
}

func main() {

	for i := 1; i <= 1000; i++ {
		result := LaunchPipeline(i)
		fmt.Println(result)
	}

}
