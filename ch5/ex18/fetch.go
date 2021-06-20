package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, url := range os.Args[1:] {
		filename, size, err := fetch(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s\t%dB\n", filename, size)
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, fmt.Errorf("failed to get %s: %v", url, err)
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "." || local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, fmt.Errorf("failed to create %s: %v", local, err)
	}
	defer func() {
		// Close file, but prefer error from Copy, if any.
		if closeErr := f.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("failed to close %s: %v", local, closeErr)
		}
	}()

	n, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", 0, fmt.Errorf("failed to copy to %s: %v", local, err)
	}

	return local, n, nil
}
