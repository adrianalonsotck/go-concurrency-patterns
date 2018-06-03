package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var timeoutMilliseconds int = 5000
var endpoints = []string{
	"https://jsonplaceholder.typicode.com/posts/1",
	"https://jsonplaceholder.typicode.com/posts/2",
	"https://jsonplaceholder.typicode.com/posts/3",
	"https://jsonplaceholder.typicode.com/posts/4",
	"https://jsonplaceholder.typicode.com/posts/5",
	"https://jsonplaceholder.typicode.com/posts/7",
	"https://jsonplaceholder.typicode.com/posts/8",
	"https://jsonplaceholder.typicode.com/posts/9",
	"https://jsonplaceholder.typicode.com/posts/10",
}

type barrierResp struct {
	Err  error
	Resp string
}

func main() {
	barrier(endpoints)
}

func barrier(endpoints []string) {
	time1 := time.Now()

	requestNumber := len(endpoints)
	in := make(chan barrierResp, requestNumber)
	defer close(in)

	responses := make([]barrierResp, requestNumber)
	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err)
			hasError = true
		}
		responses[i] = resp
	}
	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}

	time2 := time.Now()
	diff := delta(time1, time2)
	fmt.Println("Difference on seconds:", diff.Seconds())

}

func delta(time1, time2 time.Time) time.Duration {
	diferencia := time2.Sub(time1)
	return diferencia
}

func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) *
			time.Millisecond),
	}
	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	res.Resp = string(byt)
	out <- res
}
