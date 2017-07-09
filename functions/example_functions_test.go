package functions_test

import "fmt"

var (
	Echo = func(msg string) string {
		return msg
	}
	Reverse = func(msg string) string {
		last := len(msg) - 1
		var output []byte
		for i, _ := range msg {
			output = append(output, msg[last-i])
		}
		return string(output)
	}
)

func Example_functions() {
	fmt.Println(Echo("Hello"))
	fmt.Println(Reverse("Hello"))
	// Output:
	// Hello
	// olleH
}
