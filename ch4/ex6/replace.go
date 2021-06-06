package replace

import (
	"unicode"
	"unicode/utf8"
)

func ReplaceSpacesWithSingleSpace(str []byte) []byte {
	pos := 0
	spaceCount := 0

	for i := 0; i < len(str); {
		r, s := utf8.DecodeRune(str[i:])
		if !unicode.IsSpace(r) {
			if spaceCount > 0 {
				str[pos] = ' '
				pos++
			}
			copy(str[pos:], []byte(string(r)))
			pos += s
			spaceCount = 0
		} else {
			spaceCount++
		}
		i += s
	}
	if spaceCount > 0 {
		str[pos] = ' '
		pos++
	}

	return str[:pos]
}
