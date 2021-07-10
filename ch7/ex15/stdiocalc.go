package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	eval "gopl/ch7/ex13"
)

func main() {
	fmt.Printf("expression: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expr := scanner.Text()
	if err := scanner.Err(); err != nil {
		log.Fatalf("stdiocalc: %v", err)
	}

	e, err := eval.Parse(expr)
	if err != nil {
		log.Fatalf("stdiocalc: %v", err)
	}

	varset := make(map[eval.Var]bool)
	e.Check(varset)

	env := eval.Env{}
	for v := range varset {
		var value float64
		fmt.Printf("%s: ", v)
		if _, err := fmt.Scanf("%f", &value); err != nil {
			log.Fatalf("stdiocalc: %v", err)
		}
		env[v] = value
	}

	result := e.Eval(env)
	fmt.Printf("\nresult: %.6g", result)
}
