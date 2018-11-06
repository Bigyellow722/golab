package main

import (
	"fmt"
	"log"
	"sync"
	"test/comic"
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

	/*
		result, err := github.SearchIssues(os.Args[1:])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(time.Now())

		fmt.Printf("%d issues:\n", result.TotalCount)
		for _, item := range result.Items {
			fmt.Printf("create_at %v #%-5d %9.9s %.55s\n",
				item.CreatedAt, item.Number, item.User.Login, item.Title)

		}
	*/
	const rawurl string = "xkcd.com"
	client := comic.NewCrawlClient()

	var waitgroup sync.WaitGroup

	for i := 1; i < 100; i++ {
		foo := func(num int) {
			comic, err := client.Get(rawurl, num)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s: %s\n", comic.Title, comic.ImageURL)

			comic.Download(num)
			waitgroup.Done()
		}
		waitgroup.Add(1)
		go foo(i)

	}
	waitgroup.Wait()
	fmt.Println("done")
}
