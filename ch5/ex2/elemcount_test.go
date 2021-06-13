package elemcount

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestCountElements(t *testing.T) {
	input := `
		<div id="input-html">
			<div class="wrapper">
				<span><button type="button">버튼</button></span>
				<a href="https://github.com/thesulks/gopl">thesulks/gopl</a>
			</div>
		</div>
	`
	want := map[string]int{
		"div":    2,
		"span":   1,
		"a":      1,
		"button": 1,
	}

	// https://stackoverflow.com/questions/15081119/any-way-to-use-html-parse-without-it-adding-nodes-to-make-a-well-formed-tree
	// https://pkg.go.dev/golang.org/x/net/html?utm_source=gopls#ParseFragment
	nodes, err := html.ParseFragment(strings.NewReader(input), &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
	})
	if err != nil {
		t.Fatalf("failed to parse test input: %s", input)
	}

	got := make(map[string]int)
	for _, node := range nodes {
		t.Logf("%#v\n", node)
		got = CountElements(got, node)
	}

	if !equal(got, want) {
		t.Errorf("CountElements(nil, %s) = %#v, want %#v\n", input, got, want)
	}
}

func equal(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, va := range a {
		if vb, ok := b[k]; !ok || va != vb {
			return false
		}
	}
	return true
}
