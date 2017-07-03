package main

import (
	"fmt"
)

func sumItUp(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sumItUp(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sumItUp(nums...))
}
