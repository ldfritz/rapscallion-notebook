package dev

import (
	"github.com/ldfritz/rapscallion-notebook/go-api-routing/helpers"
	"io"
	"log"
	"net/http"
)

// Root is the top-level handler for this version of the API.  All
// requests for this version of the API start here.  Requests are
// directed here by `Root.ServeHTTP` in `hello.go`.  From here they will
// be forwarded onto the interesting bits.  Invalid paths raise an
// error.
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

// Greet is one of the interesting bits.  It gets a name out of the
// query string, if one was provided.  And it check the HTTP method that
// was used and adjusts its response accordingly.  Invalid methods raise
// an error.  Otherwise a brief string is returned.
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
