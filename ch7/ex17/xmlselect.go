package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", stackString(stack), tok)
			}
		}
	}
}

func stackString(s []xml.StartElement) string {
	var buf bytes.Buffer
	for _, e := range s {
		buf.WriteString(e.Name.Local)
		buf.WriteRune(' ')
	}
	return buf.String()
}

func containsAll(elems []xml.StartElement, selectors []string) bool {
	for len(selectors) <= len(elems) {
		if len(selectors) == 0 {
			return true
		}
		if contains(elems[0], selectors[0]) {
			selectors = selectors[1:]
		}
		elems = elems[1:]
	}
	return false
}

func contains(elem xml.StartElement, selector string) bool {
	if len(selector) > 0 {
		switch selector[0] {
		case '#':
			id, ok := getAttr(elem, "id")
			// log.Print(id, ok, selector)
			if !ok || id.Value != selector[1:] {
				return false
			}
			return true
		case '.':
			class, ok := getAttr(elem, "class")
			if !ok {
				return false
			}
			if classes := strings.Split(class.Value, " "); !hasElem(classes, selector[1:]) {
				return false
			}
			return true
		}
	}
	// elem
	return elem.Name.Local == selector
}

func hasElem(s []string, elem string) bool {
	for _, e := range s {
		if e == elem {
			return true
		}
	}
	return false
}

func getAttr(e xml.StartElement, attr string) (xml.Attr, bool) {
	for _, a := range e.Attr {
		// log.Print(a.Name.Local)
		if a.Name.Local == attr {
			return a, true
		}
	}
	return xml.Attr{}, false
}
