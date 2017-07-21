package main

import (
	"log"
	"strings"
)

func setup() func() {
	log.Printf(strings.Repeat("-", 20))
	return teardown
}

func teardown() {
	log.Printf(strings.Repeat("-", 20))
}

func main() {
	t := setup()
	defer t()
	log.Printf("This is my message.")
}
