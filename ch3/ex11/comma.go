package comma

import (
	"bytes"
	"strings"
)

func fComma(s string) string {
	var dot int
	if dot = strings.Index(s, "."); dot == -1 {
		dot = len(s)
	}
	if s[0] == '+' || s[0] == '-' {
		return string(s[0]) + comma(s[1:dot]) + s[dot:]
	}
	return comma(s[:dot]) + s[dot:]
}

func comma(s string) string {
	length := len(s)
	leftmost := length % 3
	if leftmost == 0 {
		leftmost = 3
	}

	buf := bytes.NewBufferString(s[:leftmost])
	for i := leftmost; i < length; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}
