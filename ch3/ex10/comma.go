package comma

import "bytes"

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
