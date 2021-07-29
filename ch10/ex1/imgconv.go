package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

var format = flag.String("format", "jpeg", "format to decode")
var supported = map[string]bool{
	"png":  true,
	"gif":  true,
	"jpeg": true,
}

func main() {
	flag.Parse()
	if !supported[*format] {
		log.Fatalf("decoder: %s not supported", *format)
	}
	toFormat(os.Stdin, os.Stdout, *format)
}

func toFormat(in io.Reader, out io.Writer, format string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch format {
	case "png":
		png.Encode(out, img)
	case "gif":
		gif.Encode(out, img, &gif.Options{NumColors: 256})
	case "jpeg":
		jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	default:
		return fmt.Errorf("toFormat: %s not supported", format)
	}
	return nil
}
