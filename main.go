package main

import (
	"os"
	"test/fetch"
)

func main() {
	for _, url := range os.Args[1:] {
		fetch.Fetch1(url)
	}
}
