package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var wg = sync.WaitGroup{}

func fetch(urls []string) {
	channel := make(chan string)
	for _, url := range urls {
		wg.Add(1)
		go getResponse(url, channel)
	}
	fmt.Println()
	for i := 0; i < len(urls); i++ {
		fmt.Printf(<-channel)
	}
	wg.Wait()
	close(channel)
}

func getResponse(url string, c chan string) {
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
	info := getContentInfo(string(body), url)
	c <- info.toStr() + "\n"
	wg.Done()
}
