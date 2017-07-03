package main

import (
	"fmt"
)

type myInt int

func (i myInt) squared() myInt {
	return i * i
}

func main() {
	num := myInt(16)
	fmt.Println(num.squared())
}
