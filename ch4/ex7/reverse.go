package reverse

import "bytes"

func utf8Reverse(bs []byte) []byte {
	rs := bytes.Runes(bs)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	bs = []byte(string(rs))
	return bs
}
