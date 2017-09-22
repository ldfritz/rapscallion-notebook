package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	filename := "output.json"

	type Person struct {
		ID   int
		Name string
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	var people []Person
	if err = json.Unmarshal(content, &people); err != nil {
		log.Printf("Unable to unmarshal content: %v", err)
	}
	for _, p := range people {
		println(p.ID, p.Name)
	}
}
