package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

type Status struct {
	ok  bool
	url string
}

func main() {
	urls := os.Args[1:]
	ch := make(chan Status, len(urls))

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, ch, &wg)
	}

	wg.Wait()
	close(ch)

	for status := range ch {
		fmt.Println(status)
	}
}

func fetch(url string, ch chan Status, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)

	if err != nil {
		ch <- Status{ok: false, url: url}
		return
	}
	defer resp.Body.Close()

	ch <- Status{ok: resp.StatusCode == http.StatusOK, url: url}
}
