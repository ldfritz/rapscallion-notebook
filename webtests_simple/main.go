package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func helloWorld(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	name := strings.Title(strings.TrimPrefix(path, "/hello/"))
	if len(name) < 1 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	io.WriteString(w, "Hello, "+name+"!\n")
}

func root(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	switch {
	case strings.HasPrefix(path, "/hello/"):
		helloWorld(w, req)
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func main() {
	port := "1234"
	http.HandleFunc("/", root)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
