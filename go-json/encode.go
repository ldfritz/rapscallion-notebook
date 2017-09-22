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

	people := []Person{
		{0, "Luke"},
		{1, "Krista"},
		{2, "Mae"},
		{3, "Dalton"},
	}

	output, err := json.Marshal(people)
	if err != nil {
		log.Printf("Unable to encode JSON: %v", err)
	}
	fmt.Println(string(output))

	if err = ioutil.WriteFile(filename, output, 0644); err != nil {
		log.Fatal(err)
	}
}
