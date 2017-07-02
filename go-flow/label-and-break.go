package main

import (
	"fmt"
)

func main() {
	var i int
	i = 1
	for i < 10 {
		for _, div := range []int{5, 7, 11} {
			if i%div == 0 {
				break
			}
		}
		i++
	}
	fmt.Println(i)

	i = 1
MainLoop:
	for i < 10 {
		for _, div := range []int{5, 7, 11} {
			if i%div == 0 {
				break MainLoop
			}
		}
		i++
	}
	fmt.Println(i)
}
