package main

import (
	"bufio"
	"fmt"
	"os"
)

func pausePrompt() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Enter line: ")
	entry, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(entry)
}

func main() {
	pausePrompt()
}
