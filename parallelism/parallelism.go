package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const MAX_PROCS = 1

var wg sync.WaitGroup

func createRandomArray(vec *[50000]int, aleatorio *rand.Rand) {
	for f := 0; f < len(vec); f++ {
		vec[f] = aleatorio.Intn(100)
	}
}

func order(vec *[50000]int) {
	for k := 1; k < len(vec); k++ {
		for f := 0; f < len(vec)-k; f++ {
			if vec[f] > vec[f+1] {
				aux := vec[f]
				vec[f] = vec[f+1]
				vec[f+1] = aux
			}
		}
	}
	wg.Done()
}

func delta(time1, time2 time.Time) time.Duration {
	diferencia := time2.Sub(time1)
	return diferencia
}

func main() {

	runtime.GOMAXPROCS(MAX_PROCS)

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	var vec1 [50000]int
	var vec2 [50000]int
	createRandomArray(&vec1, random)
	createRandomArray(&vec2, random)

	time1 := time.Now()
	wg.Add(2)
	go order(&vec1)
	go order(&vec2)
	wg.Wait()
	time2 := time.Now()

	diff := delta(time1, time2)
	fmt.Println("Difference on seconds:", diff.Seconds())
}
