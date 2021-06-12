package xkcd

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Index() error {
	latestComic, err := getLatestComic()
	if err != nil {
		return err
	}

	index := make(map[string][]int)
	for i := 1; i <= latestComic.Num; i++ {
		url := XkcdUrl + fmt.Sprintf("%d", i) + JsonPath
		comic, err := getComic(url)
		if err != nil {
			log.Printf("xkcd.Index: %v\n", err)
			continue
		}

		chucks := strings.Split(comic.Title, " ")
		for _, c := range chucks {
			c = strings.ToLower(c)
			index[c] = append(index[c], comic.Num)
		}

		var b bytes.Buffer
		encoder := gob.NewEncoder(&b)
		if err := encoder.Encode(comic); err != nil {
			return err
		}
		filename := fmt.Sprintf("comics/%d", comic.Num)
		if err := ioutil.WriteFile(filename, b.Bytes(), 0644); err != nil {
			return err
		}
	}

	var b bytes.Buffer
	encoder := gob.NewEncoder(&b)
	if err := encoder.Encode(index); err != nil {
		return err
	}
	if err := ioutil.WriteFile("index", b.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}

func getLatestComic() (*Comic, error) {
	const latestComicUrl = "https://xkcd.com/info.0.json"
	return getComic(latestComicUrl)
}

func getComic(url string) (*Comic, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed to get comic: %s", resp.Status)
	}

	var latestComic Comic
	if err := json.NewDecoder(resp.Body).Decode(&latestComic); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &latestComic, nil
}

func GetIndexInMemory() (map[string][]int, error) {
	b, err := ioutil.ReadFile("index")
	if err != nil {
		return nil, err
	}

	var index map[string][]int
	if err := gob.NewDecoder(bytes.NewBuffer(b)).Decode(&index); err != nil {
		return nil, err
	}

	return index, nil
}

func GetComic(n int) (*Comic, error) {
	filename := fmt.Sprintf("comics/%d", n)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var comic Comic
	if err := gob.NewDecoder(bytes.NewBuffer(b)).Decode(&comic); err != nil {
		return nil, err
	}

	return &comic, nil
}

func StringifyComic(c *Comic) string {
	return fmt.Sprintf("URL: %s%d\n%s\n", XkcdUrl, c.Num, c.Transcript)
}
