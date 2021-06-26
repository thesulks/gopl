package mystrings

import (
	"io/ioutil"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestNewReaderFromString(t *testing.T) {
	const filename = "golang.org.html"

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("failed to read %s: %v", filename, err)
	}
	sr := NewReaderFromString(string(bs))
	doc1, err := html.Parse(sr)
	if err != nil {
		t.Fatalf("doc1: failed to parse %s: %v", filename, err)
	}

	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("failed to open %s: %v", filename, err)
	}
	doc2, err := html.Parse(f)
	if err != nil {
		t.Fatalf("doc2: failed to parse %s: %v", filename, err)
	}

	if !naiveEqaul(doc1, doc2) {
		t.Errorf("not equal")
	}
}

func naiveEqaul(a, b *html.Node) bool {
	if a.Data != b.Data {
		return false
	}

	var ca, cb *html.Node
	for ca, cb = a.FirstChild, b.FirstChild; ca != nil && cb != nil; ca, cb = ca.NextSibling, cb.NextSibling {
		if !naiveEqaul(ca, cb) {
			return false
		}
	}

	if ca != nil || cb != nil {
		return false
	}

	return true
}
