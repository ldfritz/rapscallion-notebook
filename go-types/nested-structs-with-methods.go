package main

import (
	"fmt"
)

type Contact struct {
	firstName string
	lastName  string
}

func (c Contact) fullName() string {
	return c.firstName + " " + c.lastName
}

type Friend struct {
	Contact
	circle string
}

func (f Friend) summary() string {
	return fmt.Sprintf("%s %s", f.circle, f.fullName())
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

	for i, f := range friends {
		fmt.Printf("%02d:\n", i)
		fmt.Printf("\t%s\n", f.summary())
		fmt.Printf("\t%s %s\n", f.circle, f.fullName())
		fmt.Printf("\t%s %s %s\n", f.circle, f.firstName, f.lastName)
	}
}
