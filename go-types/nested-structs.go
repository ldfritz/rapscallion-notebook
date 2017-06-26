package main

import (
	"fmt"
)

type Contact struct {
	firstName string
	lastName  string
}

type Friend struct {
	Contact
	circle string
}

func main() {
	friends := []Friend{
		{
			Contact: Contact{
				firstName: "John",
				lastName:  "Smith",
			},
			circle: "school",
		},
		{
			Contact{
				"Jane",
				"Doe",
			},
			"work",
		},
	}

	for i, v := range friends {
		fmt.Printf("%02d: %6s %s, %s\n", i, v.circle, v.lastName, v.firstName)
	}
}
