package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(leftpad("simple", 10, '.'))
	fmt.Println(leftpad("example", 10, '-'))
	fmt.Println(leftpad("example", 3, '.'))
	fmt.Println(leftpad("example", -10, '-'))
}

func leftpad(text string, width int, char rune) string {
	if width <= len(text) {
		return text
	}
	spaces := strings.Repeat(string(char), width-len(text))
	return spaces + text
}
