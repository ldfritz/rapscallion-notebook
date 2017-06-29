package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func summarizeExistence(filename string) {
	stat, err := os.Stat(filename)
	switch {
	case err == nil:
		fmt.Printf("%s exists\n", filename)
		fmt.Println(stat)
	case os.IsNotExist(err):
		fmt.Printf("%s does not exist\n", filename)
	case err != nil:
		log.Fatal(err)
	}
}

func main() {
	filename := "tempfile.tmp"
	summarizeExistence(filename)
	fmt.Println("Writing file")

	if err := ioutil.WriteFile(filename, []byte("this is a test"), 0644); err != nil {
		log.Fatal(err)
	}
	summarizeExistence(filename)

	fmt.Println("Reading file")
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(contents))

	fmt.Println("Removing file")
	if err := os.Remove(filename); err != nil {
		log.Fatal(err)
	}
	summarizeExistence(filename)
}
