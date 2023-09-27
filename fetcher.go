package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var wg = sync.WaitGroup{}
var m = sync.Mutex{}
var results = []string{}

func fetch(urls []string) {
	for _, url := range urls {
		wg.Add(1)
		go getResponse(url)
	}
	wg.Wait()
	for _, res := range results {
		fmt.Printf(res[:100], "\n")
	}
}

func getResponse(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}
	m.Lock()
	fmt.Printf("%v done!", url)
	results = append(results, string(body))
	m.Unlock()
	wg.Done()
}
