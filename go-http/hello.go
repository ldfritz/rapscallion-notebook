package main

import (
	"io"
	"log"
	"net/http"
)

// HelloServer does something
func HelloServer(w http.ResponseWriter, req *http.Request) {
	log.Print(req)
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
