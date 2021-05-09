package main

import (
	"bufio"
	"fmt"
	"gopl/ch2/ex2/convlength"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			m, ft, err := convert(input.Text())
			if err != nil {
				fmt.Fprintf(os.Stderr, "mft: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("%s = %s, %s = %s\n", m, convlength.MToFt(m), ft, convlength.FtToM(ft))
		}
		if err := input.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "mft: %v\n", err)
			os.Exit(1)
		}
		return
	}

	for _, arg := range os.Args[1:] {
		m, ft, err := convert(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mft: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s = %s, %s = %s\n", m, convlength.MToFt(m), ft, convlength.FtToM(ft))
	}
}

func convert(s string) (convlength.Meter, convlength.Feet, error) {
	l, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, 0, err
	}
	m := convlength.Meter(l)
	ft := convlength.Feet(l)
	return m, ft, nil
}
