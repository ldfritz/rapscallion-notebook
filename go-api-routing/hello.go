package main

import (
	"github.com/ldfritz/rapscallion-notebook/go-api-routing/dev"
	"github.com/ldfritz/rapscallion-notebook/go-api-routing/v1"
	"github.com/ldfritz/rapscallion-notebook/go-api-routing/helpers"
	"log"
	"net/http"
)

// Root is a generic type for the top-level application handler instance.
type Root struct{}

// ServeHTTP on the Root will identify the initial URL branch.
func (a Root) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Print("Root.ServeHTTP(): ", req.URL.String())
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
	//	a := &Root{}
	//	log.Fatal(http.ListenAndServe(":8080", a))
	log.Fatal(http.ListenAndServe(":8080", &Root{}))
}
