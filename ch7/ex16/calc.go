package main

import (
	eval "gopl/ch7/ex13"
	"html/template"
	"log"
	"math"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("calc.html"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		expr := q.Get("expr")
		value, err := calc(expr)
		if err != nil {
			log.Printf("calc: failed to calc(%q): %v", expr, err)
			value = math.NaN()
		}
		if err := tmpl.ExecuteTemplate(w, "calc.html", value); err != nil {
			log.Printf("calc: failed to execute template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func calc(expr string) (float64, error) {
	if expr == "" {
		return 0, nil
	}
	e, err := eval.Parse(expr)
	if err != nil {
		return 0, err
	}
	// variables not allowed
	return e.Eval(eval.Env{}), nil
}
