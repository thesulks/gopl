package element

import (
	"bytes"
	"io/ioutil"
	"testing"

	"golang.org/x/net/html"
)

func TestElementByID(t *testing.T) {
	testFile := "golang.org.html"
	testHtml, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Fatalf("reading file %s failed: %v", testFile, err)
	}

	doc, err := html.Parse(bytes.NewReader(testHtml))
	if err != nil {
		t.Fatalf("parsing file %s failed: %v", testFile, err)
	}

	tests := []struct {
		input string
		want  string
	}{
		{"page", "main"},
		{"nav", "div"},
		{"test1", "span"},
		{"test2", "a"},
	}

	for _, test := range tests {
		if n := ElementByID(doc, test.input); n == nil || n.Type != html.ElementNode || n.Data != test.want {
			t.Errorf("ElementByID(%q, %q) = %q, want %q", testFile, test.input, n.Data, test.want)
		} else {
			t.Logf("ElementByID(%q, %q) = %#v", testFile, test.input, n)
		}
	}
}
