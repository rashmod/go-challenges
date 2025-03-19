package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type Site struct {
	workerId int
	url      string
	title    string
	valid    bool
	err      error
}

func main() {
	urls := os.Args[1:]

	urlChannel := make(chan string, len(urls))
	logChannel := make(chan Site, len(urls))

	var wg sync.WaitGroup

	for i := range 3 {
		wg.Add(1)
		go fetch(urlChannel, logChannel, &wg, i)
	}

	for _, url := range urls {
		urlChannel <- url
	}
	close(urlChannel)

	go func() {
		wg.Wait()
		close(logChannel)
	}()

	log(logChannel)
}

func log(ch <-chan Site) {
	for site := range ch {
		fmt.Println("worker", site.workerId, site.url)
		if site.valid {
			fmt.Println(site.title)
		} else {
			fmt.Println(site.err)
		}
		fmt.Println()
	}

}

func fetch(urlChannel <-chan string, ch chan<- Site, wg *sync.WaitGroup, i int) {
	defer wg.Done()

	for url := range urlChannel {
		resp, err := http.Get(url)
		if err != nil {
			ch <- Site{url: url, valid: false, err: err, workerId: i}
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			ch <- Site{url: url, valid: false, err: err, workerId: i}
			continue
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			ch <- Site{url: url, valid: false, err: err, workerId: i}
			continue
		}

		title := doc.Find("head title").Text()
		ch <- Site{url: url, title: title, valid: true, workerId: i}
	}
}
