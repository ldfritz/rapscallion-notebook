package main

import (
	"io"
	"log"
	"net/http"
	"path"
	"strings"
)

// ShiftPath splits off the first component of p, which will be cleaned of
// relative components before processing. `head` will never contain a slash and
// `tail` will always be a rooted path without trailing slash.
// Original: http://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html
// ***************************************************************************************
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

// ***************************************************************************************
// TODO Add a way to track a single request.
//   https://joeshaw.org/revisiting-context-and-http-handler-for-go-17/
//   https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39
// ***************************************************************************************

// AppHandler is a generic type for the top-level application handler instance.
type AppHandler struct{}

// ServeHTTP on the AppHandler will identify the initial URL branch.
func (a AppHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Print("AppHandler.ServeHTTP(): ", req.URL.String())
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	switch head {
	case "dev":
		APIHandlerv1(res, req)
	case "v1":
		APIHandlerv1(res, req)
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

// TODO Split the API versions into separate files

// APIHandlerv1 routes to the version 1 URLs.
func APIHandlerv1(res http.ResponseWriter, req *http.Request) {
	log.Print("APIHandlerv1(): ", req.URL.String())
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	switch head {
	case "hello":
		HelloHandlerv1(res, req)
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
}

// HelloHandlerv1 is the v1 hello endpoint.
func HelloHandlerv1(res http.ResponseWriter, req *http.Request) {
	log.Print("HelloHandlerv1(): ", req.URL.String())
	q := req.URL.Query()
	name := q.Get("name")
	if len(name) == 0 {
		name = "world"
	}
	io.WriteString(res, "hello, "+name+"!\n")
}

func main() {
	a := &AppHandler{}
	log.Fatal(http.ListenAndServe(":8080", a))

}
