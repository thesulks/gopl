package main

import (
	"gopl/links"
	"log"
	"net/url"
	"os"
	"strings"
)

var hosts []string

func main() {
	var rawUrls []string
	for _, inputUrl := range os.Args[1:] {
		u, err := url.Parse(inputUrl)
		if err != nil {
			log.Print(err)
			continue
		}
		hosts = append(hosts, u.Host)
		rawUrls = append(rawUrls, u.String())
	}
	log.Printf("Hosts: %v\n", hosts)

	breadthFirst(crawl, rawUrls)
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if seen[item] {
				continue
			}
			seen[item] = true
			worklist = append(worklist, f(item)...)
		}
	}
}

func crawl(rawUrl string) []string {
	// TODO: http.Get(rawUrl)의 resp.Body를 rawUrl을 경로로 하는 파일에 작성
	// fmt.Println(rawUrl)
	list, err := links.Extract(rawUrl)
	if err != nil {
		log.Print(err)
	}

	var filteredList []string
	for _, u := range list {
		if comeFromHosts(u) {
			filteredList = append(filteredList, u)
		}
	}
	return filteredList
}

func comeFromHosts(rawUrl string) bool {
	u, err := url.Parse(rawUrl)
	if err != nil {
		log.Print(err)
		return false
	}

	for _, host := range hosts {
		if strings.HasSuffix(u.Host, host) {
			return true
		}
	}

	log.Printf("%q doesn't come from target hosts\n", u.String())
	return false
}
