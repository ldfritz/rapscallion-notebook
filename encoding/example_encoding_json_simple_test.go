package encoding_test

import (
	"encoding/json"
	"fmt"
)

func Example_encodingJSONSimple() {
	type Person struct {
		Id   int
		Name string
	}
	json_data := `{"id": 1, "name": "John Doe"}`
	fmt.Println(json_data)
	person := Person{}
	json.Unmarshal([]byte(json_data), &person)
	fmt.Printf("person.Id\n// %d\n", person.Id)
	fmt.Printf("person.Name\n// %s\n", person.Name)
	// Output:
	// {"id": 1, "name": "John Doe"}
	// person.Id
	// // 1
	// person.Name
	// // John Doe
}
