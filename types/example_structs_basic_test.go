package types_test

import "fmt"

// ExampleStructs illustrates creating a simple struct.
func Example_structsBasic() {
	type contact struct {
		firstName string
		lastName  string
	}

	friends := []contact{
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
	// Output:
	// 00: Smith, John
	// 01: Doe, Jane
}
