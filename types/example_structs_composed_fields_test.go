package types_test

import "fmt"

func Example_structsComposedFields() {
	type contact struct {
		firstName string
		lastName  string
	}

	type friend struct {
		contact
		circle string
	}

	friends := []friend{
		{
			contact: contact{
				firstName: "John",
				lastName:  "Smith",
			},
			circle: "school",
		},
		{
			contact{
				"Jane",
				"Doe",
			},
			"work",
		},
	}

	for i, v := range friends {
		fmt.Printf("%02d: %6s %s, %s\n", i, v.circle, v.lastName, v.firstName)
	}
	// Output:
	// 00: school Smith, John
	// 01:   work Doe, Jane
}
