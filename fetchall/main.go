//go:build !solution

package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type res struct {
	url      string
	duration time.Duration
}

func main() {
	strings := os.Args[1:]
	start := time.Now()
	respCh := make(chan res, len(strings))
	wg := &sync.WaitGroup{}

	for _, url := range strings {
		wg.Add(1)
		go fetch(url, respCh, wg)
	}
	wg.Wait()
	close(respCh)

	for resp := range respCh {
		fmt.Printf("Url: %s\tDuration: %v\n", resp.url, resp.duration)
	}

	fmt.Println(time.Since(start))
}

func fetch(url string, respCh chan res, wg *sync.WaitGroup) {
	now := time.Now()

	resp, err := http.Get(url)
	if err == nil {
		defer resp.Body.Close()

	}

	respCh <- res{url: url, duration: time.Since(now)}
	wg.Done()
}
