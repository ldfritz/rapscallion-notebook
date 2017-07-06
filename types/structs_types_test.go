package types

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

// ExampleStructs illustrates creating a simple struct.
func ExampleStructsBasic() {
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

func ExampleStructsComposedFields() {
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

/*
func main() {
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
		fmt.Printf("\t%s\n", f.summary())
		fmt.Printf("\t%s %s\n", f.circle, f.fullName())
		fmt.Printf("\t%s %s %s\n", f.circle, f.firstName, f.lastName)
	}
}
*/

/*
package struct_test

import "testing"

func TestNamed(t *testing.T) {
	type Contact struct {
		firstName string
		lastName  string
	}

	friend := Contact{firstName: "John", lastName: "Smith"}
	t.Log(`friend := Contact{firstName: "John", lastName: "Smith"}`)
	t.Logf("friend\t// %v", friend)
	t.Logf("friend.firstName\t// %s", friend.firstName)
	t.Logf("friend.lastName\t// %s", friend.lastName)
	friend = Contact{
		firstName: "Jane",
		lastName:  "Doe",
	}
	t.Log("// Note the equals without a colon for reassignment and the mandatory trailing comma.\nfriend = Contact{\n\tfirstName: \"Jane\",\n\tlastName: \"Doe\",\n}")
	t.Logf("friend\t// %v", friend)
}

func TestPositional(t *testing.T) {
	type Contact struct {
		firstName string
		lastName  string
	}

	friend := Contact{"Jane", "Doe"}
	t.Log(`friend := Contact{"Jane", "Doe"}`)
	t.Logf("friend\t// %v", friend)
	t.Logf("friend.firstName\t// %s", friend.firstName)
	t.Logf("friend.lastName\t// %s", friend.lastName)
}

func TestMixedNamedAndPositional(t *testing.T) {
	type Contact struct {
		firstName string
		lastName  string
	}

	friends := []Contact{{firstName: "John", lastName: "Smith"}, {"Jane", "Doe"}}
	t.Log(`friends := []Contact{{firstName: "John", lastName: "Smith"}, {"Jane", "Doe"}}`)
	t.Logf("friend\t// %v", friends)
	t.Logf("friend[0].firstName\t// %s", friends[0].firstName)
	t.Logf("friend[1].lastName\t// %s", friends[1].lastName)
}

func TestChangeFields(t *testing.T) {
	type Contact struct {
		firstName string
		lastName  string
	}

	friend := Contact{"Jane", "Doe"}
	t.Log(`friend := Contact{"Jane", "Doe"}`)
	t.Logf("friend\t// %v", friend)

	friend.firstName = "John"
	t.Log(`friend.firstName = "John"`)
	t.Logf("friend\t// %v", friend)
}
*/
