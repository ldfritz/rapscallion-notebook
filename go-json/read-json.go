package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id int
	Name string
}

var json_data = `{"id": 1, "name": "John Doe"}`

func main() {
	fmt.Println(json_data)
	person := Person{}
	json.Unmarshal([]byte(json_data), &person)
	fmt.Printf("person.Id\n// %d\n", person.Id)
	fmt.Printf("person.Name\n// %s\n", person.Name)
}
