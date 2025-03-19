package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type Site struct {
	url   string
	title string
	valid bool
	err   error
}

func main() {
	urls := os.Args[1:]

	ch := make(chan Site, len(urls))

	var wg sync.WaitGroup

	fetchAll(urls, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	log(ch)
}

func log(ch chan Site) {
	for site := range ch {
		fmt.Println(site.url)
		if site.valid {
			fmt.Println(site.title)
		} else {
			fmt.Println(site.err)
		}
		fmt.Println()
	}
}

func fetchAll(urls []string, ch chan Site, wg *sync.WaitGroup) {
	wg.Add(len(urls))
	for _, url := range urls {
		go fetch(url, ch, wg)
	}
}

func fetch(url string, ch chan Site, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		ch <- Site{url: url, valid: false, err: err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		ch <- Site{url: url, valid: false, err: err}
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		ch <- Site{url: url, valid: false, err: err}
		return
	}

	title := doc.Find("head title").Text()
	ch <- Site{url: url, title: title, valid: true}
}
