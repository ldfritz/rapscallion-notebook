package interfaces_test

import "fmt"

type Modifier interface {
	Modify() string
}

func Apply(obj Modifier) string {
	return obj.Modify()
}

type Echoer struct {
	Message string
}

func (e Echoer) Modify() string {
	return e.Message
}

type Reverser struct {
	Message string
}

func (e Reverser) Modify() string {
	last := len(e.Message) - 1
	var output []byte
	for i, _ := range e.Message {
		output = append(output, e.Message[last-i])
	}
	return string(output)
}

func Example_interfaces() {
	fmt.Println(Apply(Echoer{"Hello"}))
	fmt.Println(Apply(Reverser{"Hello"}))
	// Output:
	// Hello
	// olleH
}
