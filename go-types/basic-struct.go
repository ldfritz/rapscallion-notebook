package main

import (
	"fmt"
)

type Contact struct {
	firstName string
	lastName  string
}

func main() {
	friends := []Contact{
		{
			firstName: "John",
			lastName:  "Smith",
		},
		{
			"Jane",
			"Doe",
		},
	}

	for i, v := range friends {
		fmt.Printf("%02d: %s, %s\n", i, v.lastName, v.firstName)
	}
}
