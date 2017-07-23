package main

import (
	"fmt"
	"strings"
)

func pathHeadAndTail(path string) (string, string) {
	if path == "" {
		return "", ""
	}
	slash := strings.Index(path[1:], "/") + 1
	head := path[:slash]
	tail := path[slash:]
	if slash == 0 {
		head = tail
		tail = ""
	}
	return head, tail
}

func main() {
	paths := []string{
		"one/two/three/",
		"/one/two/three/",
		"one",
		"/one",
		"",
		"/",
	}
	for _, path := range paths {
		head, tail := pathHeadAndTail(path)
		fmt.Printf("path=%s head=%s tail=%s\n", path, head, tail)
	}
}
