package main

import (
	"bufio"
	"fmt"
	"os"
)

func pausePrompt() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Press enter to continue...")
	r.ReadString('\n')
}

func main() {
	fmt.Println("This is that message that you needed to know.\nNow you know the message.\nGood day.")
	pausePrompt()
}
