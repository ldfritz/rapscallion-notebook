package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func runInteractively(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func enterToContinue(msg string) {
	r := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	_, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
}

func runQuietly(cmd *exec.Cmd) {
	var cmdout bytes.Buffer
	var cmderr bytes.Buffer
	cmd.Stdout = &cmdout
	cmd.Stderr = &cmderr
	err := cmd.Run()
	if err != nil {
		log.Print(err)
	}
	if cmdout.String() != "" {
		fmt.Print(cmdout.String())
		enterToContinue("Press enter to continue...")
	}
	if cmderr.String() != "" {
		fmt.Print(cmderr.String())
		enterToContinue("Errors returned. Press enter to continue...")
	}
}

func main() {
	enterToContinue("Press enter to launch editor")
	runInteractively(exec.Command("editor"))
	runQuietly(exec.Command("echo", "\nThe following commands will fail.\n"))
	log.Println("The following's command should fail")
	runQuietly(exec.Command("command-does-not-exist"))
	log.Println("The following's argument should fail")
	runQuietly(exec.Command("sort", "file-does-not-exist"))
}
