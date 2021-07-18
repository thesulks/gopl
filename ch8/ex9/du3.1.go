package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Message struct {
	rootId   int
	fileSize int64
}

type Root struct {
	path   string
	nfiles int64
	nbytes int64
}

var verbose = flag.Bool("v", false, "show verbose progress message")

func main() {
	flag.Parse()
	var roots []*Root
	for _, path := range flag.Args() {
		roots = append(roots, &Root{path: path})
	}
	if len(roots) == 0 {
		roots = append(roots, &Root{path: "."})
	}

	queue := make(chan *Message)
	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		go walkDir(root.path, &n, queue, i)
	}
	go func() {
		n.Wait()
		close(queue)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
Loop:
	for {
		select {
		case msg, ok := <-queue:
			if !ok {
				break Loop
			}
			id, size := msg.rootId, msg.fileSize
			roots[id].nfiles++
			roots[id].nbytes += size
		case <-tick:
			printDiskUsage(roots)
		}
	}
	printDiskUsage(roots)
}

func walkDir(dir string, n *sync.WaitGroup, queue chan<- *Message, rootId int) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, queue, rootId)
		} else {
			queue <- &Message{rootId, entry.Size()}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du3.1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(roots []*Root) {
	for _, r := range roots {
		fmt.Printf("%s: %d files  %.1f GB\n", r.path, r.nfiles, float64(r.nbytes)/1e9)
	}
}
