package main

import (
	"fmt"
	"github.com/olegnalivajev/gotuts/book_chapter1/lissajous"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main()  {
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/count", counter)
	http.HandleFunc("/liss/", liss)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request)  {
	incCount()
	_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	_, _ = fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	_, _ = fmt.Fprintf(w, "Host: %q\n", r.Host)
	_, _ = fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); r != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		_, _ = fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func incCount() {
	mu.Lock()
	count++
	mu.Unlock()
}

func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	_, _ = fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func liss(w http.ResponseWriter, r *http.Request)  {
	if err := r.ParseForm(); r != nil {
		log.Print(err)
	}
	cycles, size, delay := 5, 100, 8
	for k, v := range r.Form {
		switch k {
		case "cycles": cycles, _ = strconv.Atoi(v[0])
		case "size": size, _ = strconv.Atoi(v[0])
		case "delay": delay, _ = strconv.Atoi(v[0])
		}
	}
	lissajous.Lissajous(w, cycles, size, delay)
}