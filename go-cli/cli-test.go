package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	one   = flag.Int("one", 0, "A number.")
	two   = flag.Bool("two", false, "A flag.")
	three = flag.String("three", "", "Some text.")
)

func main() {
	fmt.Println(len(os.Args))
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}
	flag.Parse()
	fmt.Println("flag one is", *one)
	fmt.Println("flag two is", *two)
	fmt.Println("flag two is", *three)
	fmt.Println(flag.Args())
}
