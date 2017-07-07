package types_test

import "fmt"

type contact struct {
	firstName string
	lastName  string
}

type friend struct {
	contact
	circle string
}

func (c contact) fullName() string {
	return c.firstName + " " + c.lastName
}

func (f friend) summary() string {
	return fmt.Sprintf("%s %s", f.circle, f.fullName())
}

func Example_structsComposedMethods() {
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

	for i, f := range friends {
		fmt.Printf("%02d:\n", i)
		fmt.Printf("   %s\n", f.summary())
		fmt.Printf("   %s %s\n", f.circle, f.fullName())
		fmt.Printf("   %s %s %s\n", f.circle, f.firstName, f.lastName)
	}
// Output:
// 00:
//    school John Smith
//    school John Smith
//    school John Smith
// 01:
//    work Jane Doe
//    work Jane Doe
//    work Jane Doe
}
