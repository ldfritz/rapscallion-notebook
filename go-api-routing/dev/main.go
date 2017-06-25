package dev

import (
	"github.com/ldfritz/rapscallion-notebook/go-api-routing/helpers"
	"io"
	"log"
	"net/http"
)

// Root is the top-level handler for this hello version.
func Root(res http.ResponseWriter, req *http.Request) {
	log.Printf("dev.Root(): %s %s", req.Method, req.URL.String())
	var head string
	head, req.URL.Path = helpers.ShiftPath(req.URL.Path)
	switch head {
	case "hello":
		Greet(res, req)
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

// Greet is the hello endpoint.
func Greet(res http.ResponseWriter, req *http.Request) {
	log.Printf("dev.Greet(): %s %s", req.Method, req.URL.String())
	name := req.URL.Query().Get("name")
	if len(name) == 0 {
		name = "World"
	}
	switch req.Method {
	case "GET":
		io.WriteString(res, "Hello, "+name+"!\n")
	case "PUT":
		io.WriteString(res, "High-five, "+name+"!\n")
	default:
		http.Error(res, "Only GET and PUT are allowed", http.StatusMethodNotAllowed)
	}
}
