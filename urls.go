package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var urls []string = make([]string, 0, 10)

func inputs() []string {
	showList()
	listInput()
	return urls
}

func listInput() {
MainLoop:
	for {
		fmt.Print(`
		What do you want to do?
		(load) load urls from file
		(add) add urls
		(del) delete url
		(show) show the list
		(go) save and go back
		`)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		switch strings.TrimSpace(input) {
		case "load":
			load()
		case "add":
			addUrls()
		case "del":
			deleteUrl(getNumber("\nWhat is the id?\n"))
		case "show":
			showList()
		case "go":
			fmt.Print("\n")
			break MainLoop
		default:
			fmt.Print("No such option\n")
		}
	}
	save()
}

func load() {
	urltxt, err := os.ReadFile("urls.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	urls = strings.Split(string(urltxt), "\n")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, url := range urls {
		urls[i] = strings.TrimSpace(url)
	}
}

func addUrls() {
	in := ""
	for {
		in = getString("\nWhat is the url? Write \"fin\" if you're finished.\n")
		if in == "fin" {
			break
		}
		urls = append(urls, in)
	}
}

func deleteUrl(id int) {
	url := urls[id]
	if url == "" {
		fmt.Print("No such url")
		return
	}
	tmp := urls[:id]
	for i := id + 1; i < len(urls); i++ {
		tmp = append(tmp, urls[i])
	}
	urls = tmp
}

func showList() {
	for i, url := range urls {
		fmt.Printf("ID: %v	url: %v\n", i, url)
	}
	fmt.Println()
}

func save() {
	temp := []byte(strings.Join(urls, "\n"))
	err := os.WriteFile("urls.txt", temp, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func getString(q string) string {
	fmt.Print(q)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimSpace(input)
}

func getNumber(q string) int {
	fmt.Print(q)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return -1
	}
	inn, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return inn
}
