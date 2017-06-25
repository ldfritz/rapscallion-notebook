package dev

import (
	"github.com/ldfritz/rapscallion-notebook/go-api-routing/helpers"
	"io"
	"log"
	"net/http"
)

// Root is the top-level handler for this hello version.
func Root(res http.ResponseWriter, req *http.Request) {
	log.Print("dev.Root(): ", req.URL.String())
	var head string
	head, req.URL.Path = helpers.ShiftPath(req.URL.Path)
	switch head {
	case "hello":
		Greet(res, req)
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

// HelloHandlerv1 is the v1 hello endpoint.
func Greet(res http.ResponseWriter, req *http.Request) {
	log.Print("dev.Greet(): ", req.URL.String())
	name := req.URL.Query().Get("name")
	if len(name) == 0 {
		name = "World"
	}
	io.WriteString(res, "Hello, "+name+"!\n")
}
