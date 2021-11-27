package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	resp, err := fetch(os.Args[1:]...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to fetch: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Printf("GET %s\n", resp.Request.URL)

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read response body: %v\n", err)
	}
	fmt.Printf("%s", b)
}

func fetch(urls ...string) (*http.Response, error) {
	ctx, cancel := context.WithCancel(context.Background())
	reqs := make([]*http.Request, 0, len(urls))
	for _, url := range urls {
		r, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			continue
		}
		reqs = append(reqs, r)
	}

	response := make(chan *http.Response)
	wg := sync.WaitGroup{}
	errc := make(chan error)

	for _, req := range reqs {
		wg.Add(1)
		go func(r *http.Request) {
			defer wg.Done()
			resp, err := http.DefaultClient.Do(r)
			if err != nil {
				log.Printf("err: %v", err)
				errc <- err
				return
			}
			select {
			case response <- resp:
				log.Printf("%s done", r.URL.String())
			case <-ctx.Done():
				log.Printf("got %s, but already cancelled", r.URL.String())
			}
		}(req)
	}

	// error closer
	go func() {
		wg.Wait()
		close(errc)
	}()

	failed := make(chan struct{})
	// error collector
	// errs should be read only if chan failed is closed
	var errs []error
	go func() {
		for err := range errc {
			errs = append(errs, err)
		}
		close(failed)
	}()

	select {
	case resp := <-response:
		resp.Request = resp.Request.Clone(context.Background())
		cancel()
		return resp, nil
	case <-failed:
		cancel() // not really needed...
		return nil, errs[0]
	}
}
