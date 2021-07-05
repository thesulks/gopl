package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// TODO: use HTTP Method
// https://stackoverflow.com/questions/15240884/how-can-i-handle-http-requests-of-different-methods-to-in-go

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

var mu sync.Mutex
var db = database{"shoes": 50, "socks": 5}
var tmpl = template.Must(template.New("table").Parse(`
<table>
    <thead>
        <tr>
            <th>Item</th>
			<th>Price</th>
        </tr>
    </thead>
    <tbody>
		{{ range $item, $price := . }}
        <tr>
            <td>{{ $item }}</td>
            <td>{{ $price }}</td>
        </tr>
		{{ end }}
    </tbody>
</table>`))

func main() {
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/price", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	tmpl.Execute(w, db)
}

// /create?item=a&price=b
func (db database) create(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if ok {
		msg := fmt.Sprintf("already exists: %q\n", item)
		http.Error(w, msg, http.StatusConflict)
		return
	}

	price := req.URL.Query().Get("price")
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("price is invalid: %q\n", price)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	db[item] = dollars(p)
}

// /price?item=a
func (db database) read(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		msg := fmt.Sprintf("no such item: %q\n", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

// /update?item=a&price=b
func (db database) update(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		msg := fmt.Sprintf("no such item: %q\n", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	price := req.URL.Query().Get("price")
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("price is invalid: %q\n", price)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	db[item] = dollars(p)
}

// /delete?item=a
func (db database) delete(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		msg := fmt.Sprintf("no such item: %q\n", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	delete(db, item)
}
