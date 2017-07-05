package main

import "fmt"

func main() {
	data := []bool{false, true, true, false, true, true, false, false}
	var result int
	for _, v := range data {
		result = result << 1
		if v {
			result = result | 1
		}
		fmt.Printf("%b\n", result)
	}
}
