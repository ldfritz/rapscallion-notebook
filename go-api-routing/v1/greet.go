package v1

import (
	"io"
	"log"
	"net/http"
)

// Greet is one of the interesting bits.  It gets a name out of the
// query string, if one was provided.  And then it returns a string.
func Greet(res http.ResponseWriter, req *http.Request) {
	log.Printf("v1.Greet(): %s %s", req.Method, req.URL.String())
	name := req.URL.Query().Get("name")
	if len(name) == 0 {
		name = "World"
	}
	io.WriteString(res, "Hello, "+name+"!\n")
}
