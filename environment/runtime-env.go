package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Windows?")
	fmt.Println(runtime.GOOS == "windows")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
