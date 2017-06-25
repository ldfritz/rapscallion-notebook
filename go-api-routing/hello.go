package main

import (
	"github.com/ldfritz/rapscallion-notebook/go-api-routing/dev"
	"github.com/ldfritz/rapscallion-notebook/go-api-routing/helpers"
	"github.com/ldfritz/rapscallion-notebook/go-api-routing/v1"
	"log"
	"net/http"
)

// Root is a generic type for the top-level application handler instance.
type Root struct{}

// ServeHTTP on the Root will identify the initial URL branch.
func (a Root) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Printf("Root.ServeHTTP(): %s %s", req.Method, req.URL.String())
	var head string
	head, req.URL.Path = helpers.ShiftPath(req.URL.Path)
	switch head {
	case "dev":
		dev.Root(res, req)
	case "v1":
		v1.Root(res, req)
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

func main() {
	port := ":8080"
	log.Printf("Starting server on %s (^C to stop)", port)
	log.Fatal(http.ListenAndServe(port, &Root{}))
}
