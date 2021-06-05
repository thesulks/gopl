package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var bits = flag.Int("size", 256, `output size (bits):

256	SHA-256
384	SHA-384
512	SHA-512
`)

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		switch *bits {
		case 256:
			fmt.Printf("%x\n", sha256.Sum256([]byte(arg)))
		case 384:
			fmt.Printf("%x\n", sha512.Sum384([]byte(arg)))
		case 512:
			fmt.Printf("%x\n", sha512.Sum512([]byte(arg)))
		default:
			fmt.Fprintf(os.Stderr, "shash: -b=%v: invalid size\n", *bits)
		}
	}
}
