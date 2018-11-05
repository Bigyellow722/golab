package main

import (
	"fmt"
	"log"
	"os"
	"test/github"
)

func main() {
	/*
		for _, url := range os.Args[1:] {
			fetch.Fetch1(url)
		}
	*/

	/*
		ch := make(chan string)
		for _, url := range os.Args[1:] {
			go fetch.Fetch2(url, ch)
		}

		for range os.Args[1:] {
			fmt.Println(<-ch)
		}
	*/

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
