package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var urls []string
MainLoop:
	for {
		fmt.Print(`
		What do you want to do?
		(in) provide inputs
		(fet) fetch data
		(exit) exit
		`)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		switch strings.TrimSpace(input) {
		case "in":
			urls = inputs()
		case "fet":
			if len(urls) == 0 {
				println("No urls")
				break
			}
			fetch(urls)
		case "exit":
			break MainLoop
		default:
			println("No such option")
		}
	}
}
