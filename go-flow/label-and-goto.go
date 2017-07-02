package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		if i == 5 {
			goto Found5
		}
		i++
	}
Found5:
	fmt.Println(i)
}
