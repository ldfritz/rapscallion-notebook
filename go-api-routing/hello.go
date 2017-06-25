package main

import (
	"github.com/ldfritz/rapscallion-notebook/go-api-routing/dev"
	"github.com/ldfritz/rapscallion-notebook/go-api-routing/helpers"
	"github.com/ldfritz/rapscallion-notebook/go-api-routing/v1"
	"log"
	"net/http"
)

// Root is a generic type for the top-level application handler instance.
// This is given a `ServeHTTP()` method, so that it can be used by
// `http.ListenAndServe()`.
type Root struct{}

// ServeHTTP will be called for every request.  I am using it to
// identify the version for my API.  The packages are included in the
// import block.  And then requests are forwarded to the correct package
// in the switch statement.  Invalid routes throw an error.
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

// main simply announces the intent to launch the server and then does
// so on an arbitrary report.  `http.ListenAndServe` is handed a pointer
// to a Root struct which has the expected `ServeHTTP` method defined
// above.
func main() {
	port := ":8080"
	log.Printf("Starting server on %s (^C to stop)", port)
	log.Fatal(http.ListenAndServe(port, &Root{}))
}
