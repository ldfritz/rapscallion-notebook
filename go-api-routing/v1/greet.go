package v1

import (
	"io"
	"log"
	"net/http"
)

// Greet is the hello endpoint.
func Greet(res http.ResponseWriter, req *http.Request) {
	log.Printf("v1.Greet(): %s %s", req.Method, req.URL.String())
	name := req.URL.Query().Get("name")
	if len(name) == 0 {
		name = "World"
	}
	io.WriteString(res, "Hello, "+name+"!\n")
}
