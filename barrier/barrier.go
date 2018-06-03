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

type BarrierResp struct {
	Err  error
	Resp string
}

func main() {
	responses := barrier(endpoints)
	fmt.Println(responses)
}

func barrier(endpoints []string) []BarrierResp {

	requestNumber := len(endpoints)
	in := make(chan BarrierResp, requestNumber)
	defer close(in)

	responses := make([]BarrierResp, requestNumber)
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

	return responses

}

func makeRequest(out chan<- BarrierResp, url string) {
	res := BarrierResp{}
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
