package v1

import (
	"github.com/ldfritz/rapscallion-notebook/go-api-routing/helpers"
	"log"
	"net/http"
)

// Root is the top-level handler for this version of the API.  All
// requests for this version of the API start here.  Requests are
// directed here by `Root.ServeHTTP` in `hello.go`.  From here they will
// be forwarded onto the interesting bits.  Invalid paths raise an
// error.
func Root(res http.ResponseWriter, req *http.Request) {
	log.Printf("v1.Root(): %s %s", req.Method, req.URL.String())
	var head string
	head, req.URL.Path = helpers.ShiftPath(req.URL.Path)
	switch head {
	case "hello":
		Greet(res, req)
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}
