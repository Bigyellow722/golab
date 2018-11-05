package main

import (
	"fmt"
	"os"
	"test/fetch"
)

func main() {
	/*
		for _, url := range os.Args[1:] {
			fetch.Fetch1(url)
		}
	*/

	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch.Fetch2(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}
