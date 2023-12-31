package main

import (
	"fmt"
	"net/http"
)

type Engine struct{}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/index":
		http.HandleFunc(path, indexHandler)
		indexHandler(w, r)
	case "/hello":
		http.HandleFunc(path, helloHandler)
		helloHandler(w, r)
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path= %q\n", r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]= %q\n", k, v)
	}
}

func main() {
	engine := &Engine{}
	http.ListenAndServe(":8080", engine)
}
